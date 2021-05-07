package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"time"
)

func (a *Provider) DeleteInfrastructure(infra *Output) error {
	ec2Srv := ec2.New(a.session)
	elbSrv := elb.New(a.session)

	var insID []string

	for _, i := range infra.InstanceDetails {
		insID = append(insID, i.InstanceId)
	}

	err := a.DeleteInstances(insID)
	if err != nil {
		return err
	}

	if infra.LoadBalancer.LbDnsName != "" {
		_, err := elbSrv.DeleteLoadBalancer(&elb.DeleteLoadBalancerInput{LoadBalancerName: aws.String(infra.LoadBalancer.LoadBalancerName)})
		if err != nil {
			return err
		}
	}

	timeNow := time.Now()
	var nicOut *ec2.DescribeNetworkInterfacesOutput
	for time.Since(timeNow) < 5*time.Minute {

		nicOut, err = ec2Srv.DescribeNetworkInterfaces(&ec2.DescribeNetworkInterfacesInput{
			Filters: []*ec2.Filter{
				{
					Values: aws.StringSlice([]string{infra.Network.VpcId}),
					Name:   aws.String("vpc-id"),
				},
			},
		})

		if err != nil {
			return err
		}

		if len(nicOut.NetworkInterfaces) != 0 {
			continue
		}
		break
	}
	if len(nicOut.NetworkInterfaces) != 0 {
		return errors.New("some instances/lb still being used")
	}

	err = a.DeleteRoutetoRcHub(infra)
	if err != nil {
		return err
	}

	for _, sg := range infra.Network.SecurityGroups {
		_, err = ec2Srv.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
			GroupId: aws.String(sg.FirewallRuleId),
		})
		if err != nil {
			return err
		}
	}

	for _, sub := range infra.Network.Subnet {
		_, err = ec2Srv.DeleteSubnet(&ec2.DeleteSubnetInput{
			SubnetId: aws.String(sub.Id),
		})
		if err != nil {
			return err
		}
	}

	_, err = ec2Srv.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: aws.String(infra.Network.VpcId),
	})
	if err != nil {
		return err
	}

	err = a.DeleteRcInstanceProfile(infra.IAM.RoleName)
	if err != nil {
		return err
	}

	err = a.DeleteRcIamRole(infra.IAM.RoleName, infra.IAM.PolicyArn)
	if err != nil {
		return err
	}

	err = a.DeleteRcIamPolicy(infra.IAM.PolicyArn)
	if err != nil {
		return err
	}

	return nil
}
