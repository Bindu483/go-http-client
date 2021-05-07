package azure

import (
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/azure"
)

type AzureProvider struct {
	Input  Input
	Output Output
}

type ConnectionInfo struct {
	Location          string //eastus2
	ClientID          string
	ClientSecret      string
	ResourceGroupName string
	TenantID          string
	SubscriptionID    string
	Environment       *azure.Environment
}

type VNet struct {
	Name string `json:"name" yaml:"name"`
	CIDR string `json:"cidr" yaml:"cidr"`
}

type Subnet struct {
	Name string
	CIDR string //"10.0.0.0/16"
}

type NetworkInput struct {
	VNet    *VNet     `json:"vnet" yaml:"vnet"`
	Subnets []*Subnet `json:"subnet" yaml:"subnet"`
}

type NetworkSecurityGroup struct {
	NetworkSecurityGroupName string                 `json:"network_security_group_name" yaml:"network_security_group_name"`
	SecurityRule             []network.SecurityRule `json:""`
}

type LoadBalancer struct {
	Name    string `json:"name" yaml:"name"`
	PIP     string `json:"name" yaml:"name"`
	NATRule int    `json:"nat_rule_id" yaml:"nat_rule_id"`
	NIC     string `json:"nic_name" yaml:"nic_name"`
}

type Input struct {
	ConnectionInfo *ConnectionInfo
	Network        *NetworkInput
	InstanceInfo   *InstanceInfo
	LoadBalancer   *LoadBalancer
}

type InstanceInfo struct {
	InstanceName         string                `json:"instance_name" yaml:"instance_name"`
	IPName               string                `json:"ip_name" yaml:"ip_name"`
	Publisher            string                `json:"publisher" yaml:"publisher"`
	Offer                string                `json:"offer" yaml:"offer"`
	Sku                  string                `json:"sku" yaml:"sku"`
	Version              string                `json:"version" yaml:"version"`
	Username             string                `json:"username" yaml:"username"`
	SSHPublicKey         string                `json:"ssh_public_key" yaml:"ssh_public_key"`
	NIC                  string                `json:"nic_name" yaml:"nic_name"`
	NetworkSecurityGroup *NetworkSecurityGroup `json:"network_security_group" yaml:"network_security_group"`
	AvailabilityGroup    string                `json:"availability_group" yaml:"availability_group"`
}

type Output struct {
	SSHPrivateKey string
}
