package aws

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
)

func (a *Provider) CreateFirewallRule(vpcId string, rules *cloud_providers.FireWallRulesInput, tags AWSTags) (firewallOutput *cloud_providers.FireWallRuleOutput, err error) {
	ec2Srv := ec2.New(a.session)

	sgOut, err := ec2Srv.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		VpcId:       aws.String(vpcId),
		GroupName:   aws.String(rules.Name),
		Description: aws.String(rules.Name),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create security group %s err %s", rules.Name, err.Error()))
	}

	updatedTags := tags
	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = rules.Name
	}
	err = addEc2Tags(a.getTags(updatedTags), ec2Srv, []*string{sgOut.GroupId})
	if err != nil {
		return nil, err
	}

	for _, r := range rules.Rules {
		if r.Type == cloud_providers.IngressRule {
			err := AddInboundSgRule(ec2Srv, r, sgOut.GroupId)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("failed to add inbound rule to sg %s", *sgOut.GroupId))
			}
		} else if r.Type == cloud_providers.EgressRule {
			err := AddOutboundSgRule(ec2Srv, r, sgOut.GroupId)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("failed to add inbound rule to sg %s", *sgOut.GroupId))
			}
		}
	}

	return &cloud_providers.FireWallRuleOutput{Name: rules.Name, FirewallRuleId: aws.StringValue(sgOut.GroupId)}, nil
}

func AddInboundSgRule(e *ec2.EC2, rule *cloud_providers.FirewallRule, sgId *string) error {
	_, err := e.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(rule.FromPort),
				ToPort:     aws.Int64(rule.ToPort),
				IpProtocol: aws.String(rule.Protocol),
				IpRanges:   []*ec2.IpRange{{CidrIp: aws.String(rule.FromCIDR)}},
			},
		},
		GroupId: sgId,
	})
	if err != nil {
		return err
	}
	return nil
}

func AddOutboundSgRule(e *ec2.EC2, rule *cloud_providers.FirewallRule, sgId *string) error {

	_, err := e.AuthorizeSecurityGroupEgress(&ec2.AuthorizeSecurityGroupEgressInput{
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(rule.FromPort),
				ToPort:     aws.Int64(rule.ToPort),
				IpProtocol: aws.String(rule.Protocol),
				IpRanges:   []*ec2.IpRange{{CidrIp: aws.String(rule.FromCIDR)}},
			},
		},
		GroupId: sgId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *Provider) CreateVPC(cidr string, tags AWSTags) (string, error) {
	ec2Srv := ec2.New(a.session)
	v, err := ec2Srv.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:       aws.String(cidr),
		InstanceTenancy: aws.String("default"),
	})
	if err != nil {
		return "", errors.New("failed to create vpc " + err.Error())
	}

	_, err = ec2Srv.ModifyVpcAttribute(&ec2.ModifyVpcAttributeInput{
		VpcId:            v.Vpc.VpcId,
		EnableDnsSupport: &ec2.AttributeBooleanValue{Value: aws.Bool(true)}})
	if err != nil {
		return "", errors.New("failed to enable dns support for vpc " + err.Error())
	}

	updatedTags := tags
	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = fmt.Sprintf("rc-vpc-%s", a.DeploymentName)
	}
	err = addEc2Tags(a.getTags(updatedTags), ec2Srv, []*string{v.Vpc.VpcId})
	if err != nil {
		return "", err
	}

	return *v.Vpc.VpcId, nil
}

func (a *Provider) CreateSubnet(vpcId string, cidr, name string, availabilityZone string, tags AWSTags) (string, error) {
	ec2Srv := ec2.New(a.session)
	s, err := ec2Srv.CreateSubnet(&ec2.CreateSubnetInput{
		VpcId:            aws.String(vpcId),
		CidrBlock:        aws.String(cidr),
		AvailabilityZone: aws.String(availabilityZone),
	})
	if err != nil {
		return "", errors.New("failed to create subnet " + err.Error())
	}

	_, err = ec2Srv.ModifySubnetAttribute(&ec2.ModifySubnetAttributeInput{
		MapPublicIpOnLaunch: &ec2.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
		SubnetId: s.Subnet.SubnetId,
	})
	if err != nil {
		return *s.Subnet.SubnetId, errors.New("failed to enable public ip on subnet " + err.Error())
	}

	updatedTags := tags

	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = name
	}
	err = addEc2Tags(a.getTags(updatedTags), ec2Srv, []*string{s.Subnet.SubnetId})
	if err != nil {
		return "", err
	}
	return *s.Subnet.SubnetId, nil
}

func (a *Provider) DeleteRoutetoRcHub(output *Output) (err error) {

	ec2Srv := ec2.New(a.session)

	//Disassociate the Subnet from Route table

	routeTableInfo, err := ec2Srv.DescribeRouteTables(&ec2.DescribeRouteTablesInput{Filters: []*ec2.Filter{
		{
			Name:   aws.String("vpc-id"),
			Values: []*string{aws.String(output.Network.VpcId)},
		},
	}})
	if err != nil {
		return err
	}

	if routeTableInfo == nil {
		return errors.New("error fetching route table info")
	}

	for _, routeTable := range routeTableInfo.RouteTables {
		for _, associationID := range routeTable.Associations {
			_, err = ec2Srv.DisassociateRouteTable(&ec2.DisassociateRouteTableInput{AssociationId: associationID.RouteTableAssociationId})
			if err != nil {
				continue
			}
			_, err = ec2Srv.DeleteRouteTable(&ec2.DeleteRouteTableInput{RouteTableId: routeTable.RouteTableId})
			if err != nil {
				continue
			}
		}
	}

	_, err = ec2Srv.DetachInternetGateway(&ec2.DetachInternetGatewayInput{
		InternetGatewayId: aws.String(output.Network.InternetGatewayId),
		VpcId:             aws.String(output.Network.VpcId),
	})
	if err != nil {
		return err
	}

	_, err = ec2Srv.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{InternetGatewayId: aws.String(output.Network.InternetGatewayId)})
	if err != nil {
		return err
	}

	return nil
}

func (a *Provider) CreateRouteToRcHub(vpcId string, subnetId []string, tags AWSTags) (igwId string, rtId string, err error) {
	ec2Srv := ec2.New(a.session)
	igwOut, err := ec2Srv.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	if err != nil {
		return "", "", errors.New("failed to create internet gateway " + err.Error())
	}

	_, err = ec2Srv.AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		VpcId:             aws.String(vpcId),
		InternetGatewayId: igwOut.InternetGateway.InternetGatewayId,
	})
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to associate internet gateway with vpc %s", err.Error()))
	}

	updatedTags := tags
	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = fmt.Sprintf("rc-igw-%s", a.DeploymentName)
	}
	err = addEc2Tags(a.getTags(updatedTags), ec2Srv, []*string{igwOut.InternetGateway.InternetGatewayId})
	if err != nil {
		return "", "", err
	}

	rTable, err := ec2Srv.CreateRouteTable(&ec2.CreateRouteTableInput{VpcId: aws.String(vpcId)})
	if err != nil {
		return "", "", errors.New("failed to create route table " + err.Error())
	}

	updatedTags["Name"] = fmt.Sprintf("rc-rt-main-%s", a.DeploymentName)
	err = addEc2Tags(updatedTags, ec2Srv, []*string{rTable.RouteTable.RouteTableId})
	if err != nil {
		return "", "",
			errors.New(fmt.Sprintf("failed to add route to rt %s err %s", *rTable.RouteTable.RouteTableId, err.Error()))
	}

	_, err = ec2Srv.CreateRoute(&ec2.CreateRouteInput{
		RouteTableId:         rTable.RouteTable.RouteTableId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            igwOut.InternetGateway.InternetGatewayId,
	})
	if err != nil {
		return "", "", errors.New("failed to add route to internet in the route table")
	}

	for _, s := range subnetId {
		_, err = ec2Srv.AssociateRouteTable(&ec2.AssociateRouteTableInput{
			RouteTableId: rTable.RouteTable.RouteTableId,
			SubnetId:     aws.String(s),
		})
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to associate rt with subnet %s", err.Error()))
		}
	}

	return *igwOut.InternetGateway.InternetGatewayId, *rTable.RouteTable.RouteTableId, nil
}
