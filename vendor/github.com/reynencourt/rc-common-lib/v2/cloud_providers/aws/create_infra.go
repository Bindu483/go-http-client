package aws

import (
	"errors"
	"fmt"

	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
)

func (a *Provider) CreateInfrastructure(input *Input) (*Output, error) {

	//Network stack
	network := NetworkOutput{}
	vpcId, err := a.CreateVPC(input.Network.VpcCidr, nil)
	if err != nil {
		return nil, err
	}

	network.VpcId = vpcId

	subnetOutput := make([]SubnetOutput, len(input.Network.Subnet))
	subnetIds := make([]string, len(input.Network.Subnet))
	if len(input.Network.Subnet) < 2 {
		return nil, errors.New("subnet count should be minimum two.")
	}
	for i, s := range input.Network.Subnet {
		zone := "a"
		if i%2 == 0 {
			zone = "b"
		}
		subnet, err := a.CreateSubnet(vpcId, s.Cidr, s.Name, fmt.Sprintf("%s%s", input.Region, zone), nil)
		if err != nil {
			return nil, err
		}
		subnetOutput[i] = SubnetOutput{Name: s.Name, Id: subnet}
		subnetIds[i] = subnet
	}

	network.Subnet = subnetOutput

	igwId, rtId, err := a.CreateRouteToRcHub(vpcId, subnetIds, nil)
	if err != nil {
		return nil, err
	}

	network.RouteTableId = rtId
	network.InternetGatewayId = igwId

	var sgOut = make([]cloud_providers.FireWallRuleOutput, len(input.Network.SecurityGroups))
	for i, sg := range input.Network.SecurityGroups {
		out, err := a.CreateFirewallRule(vpcId, &sg, nil)
		if err != nil {
			return nil, err
		}
		sgOut[i] = *out
	}
	network.SecurityGroups = sgOut

	//IAM stack
	iam := IAMOutput{}
	iamPolicyArn, err := a.CreateRcIamPolicy(input.Iam.PolicyName)
	if err != nil {
		return nil, err
	}
	iam.PolicyArn = iamPolicyArn

	iamRoleArn, err := a.CreateRcIamRole(input.Iam.RoleName, iamPolicyArn, nil)
	if err != nil {
		return nil, err
	}
	iam.RoleArn = iamRoleArn
	iam.RoleName = input.Iam.RoleName

	iamProfileArn, err := a.CreateRcInstanceProfile(input.Iam.RoleName)
	if err != nil {
		return nil, err
	}
	iam.ProfileArn = iamProfileArn
	iam.ProfileName = fmt.Sprintf("%s-profile", input.Iam.RoleName)

	//ec2 stack

	instanceSgIds := make([]string, len(input.InstanceInfo.SgName))
	for i, s := range input.InstanceInfo.SgName {
		instanceSgIds[i] = network.SecurityGroups.GetFirewalId(s)
	}
	req := &CreateInstanceRequest{
		SubnetId:              subnetIds[0],
		IamInstanceProfileArn: iamProfileArn,
		OsType:                input.InstanceInfo.OsType,
		InstanceType:          input.InstanceInfo.InstanceSize,
		KeyName:               input.InstanceInfo.KeyName,
		SecurityGroupIds:      instanceSgIds,
		Count:                 1,
		DiskSize:              100,
	}

	instance, err := a.CreateInstances(*req)
	if err != nil {
		return nil, err
	}

	// loadbalancer stack
	var lb LoadBalancerOutputDetails
	if input.LoadBalancer.Required {
		lb = LoadBalancerOutputDetails{}
		lbSgIds := make([]string, len(input.LoadBalancer.SgNames))
		for i, s := range input.LoadBalancer.SgNames {
			lbSgIds[i] = network.SecurityGroups.GetFirewalId(s)
		}
		subnetIds := make([]string, len(input.LoadBalancer.SubnetNames))
		for i, s := range input.LoadBalancer.SubnetNames {
			subnetIds[i] = network.Subnet.GetSubnetId(s)
		}
		var scheme LBScheme
		if input.LoadBalancer.IsPublic {
			scheme = LBSchemePublic
		} else {
			scheme = LBSchemePrivate
		}

		lbReq := &ClassicLbRequest{
			Tags:              nil,
			SecurityGroupIds:  lbSgIds,
			SubnetId:          subnetIds,
			Name:              fmt.Sprintf("rcp-lb-%s", a.DeploymentName),
			InstancePort:      80,
			LbPort:            80,
			Protocol:          HTTP,
			Scheme:            scheme,
			InstancesToAttach: instance.GetInstanceIds(),
		}

		lbDnsName, err := a.CreateClassicLoadBalancer(lbReq)
		if err != nil {
			return nil, err
		}
		lb.LbDnsName = lbDnsName
		lb.LoadBalancerName = fmt.Sprintf("rcp-lb-%s", a.DeploymentName)
	}

	var ins []Instance
	for _, i := range instance {
		ins = append(ins, *i)
	}

	return &Output{IAM: iam, Network: network, InstanceDetails: ins, LoadBalancer: lb}, nil
}
