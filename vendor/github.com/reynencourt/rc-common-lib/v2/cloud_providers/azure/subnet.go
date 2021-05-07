package azure

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"strings"
)

func (a *AzureProvider) getSubnetsClient() (*network.SubnetsClient, error) {
	subnetsClient := network.NewSubnetsClient(a.Input.ConnectionInfo.SubscriptionID)

	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	subnetsClient.Authorizer = auth

	return &subnetsClient, nil
}

func (a *AzureProvider) validateVPC() error {

	if strings.TrimSpace(a.Input.Network.VNet.Name) == "" {
		return errors.New("vpc name cannot be empty")
	}

	if len(a.Input.Network.Subnets) == 0 {
		return errors.New("atleast one subnet is required")
	}

	if strings.TrimSpace(a.Input.ConnectionInfo.ResourceGroupName) == "" {
		return errors.New("resource name cannot be empty")
	}

	return nil
}

func (a *AzureProvider) GetVirtualNetworkSubnet(ctx context.Context) (*network.Subnet, error) {

	if err := a.validateVPC(); err != nil {
		return nil, err
	}

	subnetsClient, err := a.getSubnetsClient()
	if err != nil {
		return nil, err
	}

	subnet, err := subnetsClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.Network.VNet.Name, a.Input.Network.Subnets[0].Name, "")
	if err != nil {
		return nil, err
	}

	return &subnet, nil
}
