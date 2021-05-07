package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

func (a *AzureProvider) getToken() (*autorest.BearerAuthorizer, error) {
	oauthConfig, err := adal.NewOAuthConfig(
		a.Input.ConnectionInfo.Environment.ActiveDirectoryEndpoint, a.Input.ConnectionInfo.TenantID)
	if err != nil {
		return nil, err
	}

	token, err := adal.NewServicePrincipalToken(
		*oauthConfig,
		a.Input.ConnectionInfo.ClientID,
		a.Input.ConnectionInfo.ClientSecret,
		a.Input.ConnectionInfo.Environment.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	return autorest.NewBearerAuthorizer(token), nil
}

func (a *AzureProvider) SetNetwork(network *NetworkInput) *AzureProvider {
	a.Input.Network = network
	return a
}

func (a *AzureProvider) SetInstanceInfo(instanceInfo *InstanceInfo) *AzureProvider {
	a.Input.InstanceInfo = instanceInfo
	return a
}

func (a *AzureProvider) SetLoadBalancer(loadBalancer *LoadBalancer) *AzureProvider {
	a.Input.LoadBalancer = loadBalancer
	return a
}

func (this *NetworkInput) SetVNet(vnet *VNet) *NetworkInput {
	this.VNet = vnet
	return this
}

func (this *NetworkInput) SetSubnet(subnets []*Subnet) *NetworkInput {
	this.Subnets = subnets
	return this
}

func (this *InstanceInfo) SetInstanceNetworkSecurityGroups(networkSecurityGroups *NetworkSecurityGroup) *InstanceInfo {
	this.NetworkSecurityGroup = networkSecurityGroups
	return this
}

func NewWithConnection(connection *ConnectionInfo) (*AzureProvider, error) {

	azureEnv, err := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail
	if err != nil {
		return nil, err
	}

	connection.Environment = &azureEnv

	var azureProvider AzureProvider

	input := Input{
		ConnectionInfo: connection,
	}

	azureProvider.Input = input
	return &azureProvider, nil
}

func New(location string,
	tenantID string,
	subscriptionId string,
	clientID string,
	clientSecret string,
	resourceGroupName string) (*AzureProvider, error) {

	azureEnv, err := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail
	if err != nil {
		return nil, err
	}

	var azureProvider AzureProvider

	input := Input{
		ConnectionInfo: &ConnectionInfo{
			Location:          location,
			ResourceGroupName: resourceGroupName,
			SubscriptionID:    subscriptionId,
			Environment:       &azureEnv,
			TenantID:          tenantID,
			ClientSecret:      clientSecret,
			ClientID:          clientID,
		},
	}

	azureProvider.Input = input
	return &azureProvider, nil
}
