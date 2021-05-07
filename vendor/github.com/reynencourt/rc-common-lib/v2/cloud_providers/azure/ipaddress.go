package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func (a *AzureProvider) CreatePublicIP(ctx context.Context, ipName string) (*network.PublicIPAddress, error) {
	ipClient, err := a.getIPClient()
	if err != nil {
		return nil, err
	}

	future, err := ipClient.CreateOrUpdate(
		ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		ipName,
		network.PublicIPAddress{
			Name:     to.StringPtr(ipName),
			Location: to.StringPtr(a.Input.ConnectionInfo.Location),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAddressVersion:   network.IPv4,
				PublicIPAllocationMethod: network.Static,
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("cannot create public ip address: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, ipClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get public ip address create or update future response: %v", err)
	}

	ip, err := future.Result(*ipClient)
	if err != nil {
		return nil, fmt.Errorf("cannot get public ip address create or update future response: %v", err)
	}

	return &ip, nil
}

// GetPublicIP returns an existing public IP
func (a *AzureProvider) GetPublicIP(ctx context.Context, ipName string) (*network.PublicIPAddress, error) {
	ipClient, err := a.getIPClient()
	if err != nil {
		return nil, err
	}

	ip, err := ipClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, ipName, "")
	if err != nil {
		return nil, err
	}

	return &ip, nil
}

// DeletePublicIP deletes an existing public IP
func (a *AzureProvider) DeletePublicIP(ctx context.Context, ipName string) (result network.PublicIPAddressesDeleteFuture, err error) {
	ipClient, err := a.getIPClient()
	if err != nil {
		return
	}
	return ipClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName, ipName)
}

func (a *AzureProvider) getIPClient() (*network.PublicIPAddressesClient, error) {
	ipClient := network.NewPublicIPAddressesClient(a.Input.ConnectionInfo.SubscriptionID)
	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	ipClient.Authorizer = auth

	return &ipClient, nil
}
