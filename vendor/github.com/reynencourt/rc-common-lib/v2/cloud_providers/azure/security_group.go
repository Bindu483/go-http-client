package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func (a *AzureProvider) getNsgClient() (*network.SecurityGroupsClient, error) {
	nsgClient := network.NewSecurityGroupsClient(a.Input.ConnectionInfo.SubscriptionID)
	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	nsgClient.Authorizer = auth
	return &nsgClient, nil
}

// GetNetworkSecurityGroup returns an existing network security group
func (a *AzureProvider) GetInstanceSecurityGroup(ctx context.Context) (sg network.SecurityGroup, err error) {
	nsgClient, err := a.getNsgClient()
	if err != nil {
		return
	}

	return nsgClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.InstanceInfo.NetworkSecurityGroup.NetworkSecurityGroupName, "")
}

func (a *AzureProvider) CreateInstanceSecurityGroup(ctx context.Context) (*network.SecurityGroup, error) {
	nsgClient, err := a.getNsgClient()
	if err != nil {
		return nil, err
	}

	future, err := nsgClient.CreateOrUpdate(
		ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		a.Input.InstanceInfo.NetworkSecurityGroup.NetworkSecurityGroupName,
		network.SecurityGroup{
			Location: to.StringPtr(a.Input.ConnectionInfo.Location),
			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
				SecurityRules: &a.Input.InstanceInfo.NetworkSecurityGroup.SecurityRule,
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("cannot create nsg: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, nsgClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get nsg create or update future response: %v", err)
	}

	nsg, err := future.Result(*nsgClient)
	if err != nil {
		return nil, fmt.Errorf("cannot get nsg create or update future response: %v", err)
	}

	return &nsg, nil
}
