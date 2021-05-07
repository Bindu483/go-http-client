package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
)

type Input struct {
	Region       string                   `json:"region" yaml:"region"`
	Network      NetworkInput             `json:"network" yaml:"network"`
	LoadBalancer LoadBalancerInputDetails `json:"load_balancer" yaml:"load_balancer"`
	InstanceInfo InstanceDetailsInput     `json:"instance_info" yaml:"instance_info"`
	Iam          IamInput                 `json:"iam" yaml:"iam"`
}

type NetworkInput struct {
	VpcCidr        string                               `json:"vpc_cidr" yaml:"vpc_cidr"`
	Subnet         []SubnetInput                        `json:"subnet" yaml:"subnet"`
	WhiteListedIps []string                             `json:"white_listed_ips" yaml:"white_listed_ips"`
	SecurityGroups []cloud_providers.FireWallRulesInput `json:"security_groups" yaml:"security_groups"`
}

type SubnetInput struct {
	Name string `json:"name" yaml:"name"`
	Cidr string `json:"cidr" yaml:"cidr"`
}
type SubnetOutput struct {
	Name string `json:"name" yaml:"name"`
	Id   string `json:"id" yaml:"id"`
}

type InstanceDetailsInput struct {
	InstanceSize   string                 `json:"instance_size" yaml:"instance_size"`
	OsType         cloud_providers.OsType `json:"os_type" yaml:"os_type"`
	KeyName        string                 `json:"key_name" yaml:"key_name"`
	KeyPath        string                 `json:"key_path" yaml:"key_path"`
	SgName         []string               `json:"sg_name" yaml:"sg_name"`
	IamProfileName string                 `json:"iam_profile_name" yaml:"iam_profile_name"`
	SubnetName     string                 `json:"subnet_name" yaml:"subnet_name"`
}

type LoadBalancerInputDetails struct {
	Required    bool     `json:"required" yaml:"required"`
	IsPublic    bool     `json:"is_public" yaml:"is_public"`
	SgNames     []string `json:"sg_names" yaml:"sg_names"`
	SubnetNames []string `json:"subnet_names" yaml:"subnet_names"`
	Name        string   `json:"name" yaml:"name"`
}

type NetworkOutput struct {
	VpcId             string              `json:"vpc_id" yaml:"vpc_id"`
	RouteTableId      string              `json:"route_table_id" yaml:"route_table_id"`
	InternetGatewayId string              `json:"internet_gateway_id" yaml:"internet_gateway_id"`
	Subnet            Subnets             `json:"subnet" yaml:"subnet"`
	SecurityGroups    FireWallRuleOutputs `json:"security_groups" yaml:"security_groups"`
}

type Subnets []SubnetOutput

type FireWallRuleOutputs []cloud_providers.FireWallRuleOutput

type IAMOutput struct {
	PolicyArn   string `json:"policy_arn" yaml:"policy_arn"`
	RoleArn     string `json:"role_arn" yaml:"role_arn"`
	RoleName    string `json:"role_name" yaml:"role_name"`
	ProfileArn  string `json:"profile_arn" yaml:"profile_arn"`
	ProfileName string `json:"profile_name" yaml:"profile_name"`
}

type IamInput struct {
	RoleName   string `json:"role_name" yaml:"role_name"`
	PolicyName string `json:"policy_name" yaml:"policy_name"`
}

type LoadBalancerOutputDetails struct {
	LoadBalancerName string `json:"load_balancer_name" yaml:"load_balancer_name"`
	LbDnsName        string `json:"lb_dns_name" yaml:"lb_dns_name"`
}

type Output struct {
	Network         NetworkOutput             `json:"network" yaml:"network"`
	InstanceDetails []Instance                `json:"instance_details" yaml:"instance_details"`
	IAM             IAMOutput                 `json:"iam" yaml:"iam"`
	LoadBalancer    LoadBalancerOutputDetails `json:"load_balancer" yaml:"load_balancer"`
}

type Instances []*Instance

type Instance struct {
	Name            string   `json:"name" yaml:"name"`
	PublicIp        string   `json:"public_ip" yaml:"public_ip"`
	PrivateIp       string   `json:"private_ip" yaml:"private_ip"`
	InstanceId      string   `json:"instance_id" yaml:"instance_id"`
	SshKeyName      string   `json:"ssh_key_name" yaml:"ssh_key_name"`
	SSHPublicKey    string   `json:"ssh_public_key" yaml:"ssh_public_key"`
	SSHPrivateKey   string   `json:"ssh_private_key" yaml:"ssh_private_key"`
	SecurityGroupId []string `json:"node_security_group_id" yaml:"node_security_group_id"`
}

type Provider struct {
	session        *session.Session
	DeploymentName string
	ConsumerKey    string
	Region         string
}

type Infrastructure struct {
	Input  *Input  `json:"input" yaml:"input"`
	Output *Output `json:"output" yaml:"output"`
}
