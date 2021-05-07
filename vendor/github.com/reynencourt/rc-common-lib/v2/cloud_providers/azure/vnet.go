package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
)

//see https://github.com/Azure-Samples/azure-sdk-for-go-samples/commit/4c0c9a00d0914bf9cfaf251c257eac9d75482f6b
func (a *AzureProvider) getVnetClient() (*network.VirtualNetworksClient, error) {
	vnetClient := network.NewVirtualNetworksClient(a.Input.ConnectionInfo.SubscriptionID)

	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	vnetClient.Authorizer = auth

	return &vnetClient, nil
}

// CreateVirtualNetworkAndSubnets creates a virtual network with two subnets
// Replace with input
func (a *AzureProvider) CreateVirtualNetworkAndSubnets(ctx context.Context) (*network.VirtualNetwork, error) {

	if err := a.validateVPC(); err != nil {
		return nil, err
	}

	vnetClient, err := a.getVnetClient()
	if err != nil {
		return nil, err
	}

	var subnetInfo []network.Subnet
	for _, subnet := range a.Input.Network.Subnets {
		subnetInfo = append(subnetInfo, network.Subnet{
			Name: to.StringPtr(subnet.Name),
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr(subnet.CIDR),
			},
		})
	}

	future, err := vnetClient.CreateOrUpdate(
		ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		a.Input.Network.VNet.Name,
		network.VirtualNetwork{
			Location: to.StringPtr(a.Input.ConnectionInfo.Location),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{a.Input.Network.VNet.CIDR},
				},
				Subnets: &subnetInfo,
			},
		})

	if err != nil {
		return nil, fmt.Errorf("cannot create virtual network: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vnetClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the vnet create or update future response: %v", err)
	}

	res, err := future.Result(*vnetClient)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteVirtualNetwork deletes a virtual network given an existing virtual network
func (a *AzureProvider) DeleteVirtualNetwork(ctx context.Context) (result network.VirtualNetworksDeleteFuture, err error) {
	vnetClient, err := a.getVnetClient()
	return vnetClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.Network.VNet.Name)
}
