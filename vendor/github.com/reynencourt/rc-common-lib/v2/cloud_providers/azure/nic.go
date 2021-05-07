package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"log"
)

func (a *AzureProvider) getNicClient() (*network.InterfacesClient, error) {
	nicClient := network.NewInterfacesClient(a.Input.ConnectionInfo.SubscriptionID)

	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	nicClient.Authorizer = auth
	return &nicClient, nil
}

// GetNic returns an existing network interface
func (a *AzureProvider) GetLoadBalncerNic(ctx context.Context) (n network.Interface, err error) {
	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	return nicClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.LoadBalancer.NIC, "")
}

func (a *AzureProvider) GetClusterNic(ctx context.Context, nicName string) (*network.Interface, error) {
	nicClient, err := a.getNicClient()
	if err != nil {
		return nil, err
	}

	nic, err := nicClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, nicName, "")
	if err != nil {
		return nil, err
	}

	return &nic, nil
}

// GetNic returns an existing network interface
func (a *AzureProvider) GetInstanceNic(ctx context.Context) (n network.Interface, err error) {
	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	return nicClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.InstanceInfo.NIC, "")
}

// DeleteNic deletes an existing network interface
func (a *AzureProvider) DeleteLoadBalancerNic(ctx context.Context) (result network.InterfacesDeleteFuture, err error) {
	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	return nicClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.LoadBalancer.NIC)
}

// DeleteNic deletes an existing network interface
func (a *AzureProvider) DeleteInstanceNic(ctx context.Context) (result network.InterfacesDeleteFuture, err error) {
	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	return nicClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.InstanceInfo.NIC)
}

// CreateNICWithLoadBalancer creats a network interface, wich is set up with a loadbalancer's NAT rule
func (a *AzureProvider) CreateNICWithLoadBalancer(ctx context.Context) (nic network.Interface, err error) {
	subnet, err := a.GetVirtualNetworkSubnet(ctx)
	if err != nil {
		return
	}

	lb, err := a.CreateLoadBalancer(ctx)
	if err != nil {
		return
	}

	conn := a.Input.ConnectionInfo
	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	future, err := nicClient.CreateOrUpdate(ctx,
		conn.ResourceGroupName,
		a.Input.LoadBalancer.NIC,
		network.Interface{
			Location: to.StringPtr(conn.Location),
			InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
				IPConfigurations: &[]network.InterfaceIPConfiguration{
					{
						Name: to.StringPtr("pipConfig"),
						InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
							Subnet: &network.Subnet{
								ID: subnet.ID,
							},
							LoadBalancerBackendAddressPools: &[]network.BackendAddressPool{
								{
									ID: (*lb.BackendAddressPools)[0].ID,
								},
							},
						},
					},
				},
			},
		})
	if err != nil {
		return nic, fmt.Errorf("cannot create nic: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, nicClient.Client)
	if err != nil {
		return nic, fmt.Errorf("cannot get nic create or update future response: %v", err)
	}

	return future.Result(*nicClient)
}

func (a *AzureProvider) CreateClusterNic(ctx context.Context, instanceName string) (i network.Interface, err error) {
	subnet, err := a.GetVirtualNetworkSubnet(ctx)
	if err != nil {
		return
	}

	nicParams := network.Interface{
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			IPConfigurations: &[]network.InterfaceIPConfiguration{
				{
					Name: to.StringPtr(instanceName),
					InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
						Subnet:                    subnet,
						PrivateIPAllocationMethod: network.Dynamic,
					},
				},
			},
		},
	}

	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	future, err := nicClient.CreateOrUpdate(ctx, a.Input.ConnectionInfo.ResourceGroupName, instanceName, nicParams)
	if err != nil {
		return
	}

	err = future.WaitForCompletionRef(ctx, nicClient.Client)
	if err != nil {
		return
	}

	return future.Result(*nicClient)

}

func (a *AzureProvider) CreateInstanceNIC(ctx context.Context, ipName string) (i network.Interface, err error) {

	subnet, err := a.GetVirtualNetworkSubnet(ctx)
	if err != nil {
		log.Fatalf("failed to get subnet: %v", err)
	}

	ip, err := a.GetPublicIP(ctx, ipName)
	if err != nil {
		log.Fatalf("failed to get ip address: %v", err)
	}

	var pools *[]network.BackendAddressPool
	var lb *network.LoadBalancer

	if a.Input.LoadBalancer != nil {

		lb, err = a.GetLoadBalancer(ctx)
		if err != nil {
			return i, err
		}

		pools = &[]network.BackendAddressPool{
			{
				ID: (*lb.BackendAddressPools)[0].ID,
			},
		}
	}

	nicParams := network.Interface{
		Name:     to.StringPtr(a.Input.InstanceInfo.NIC),
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			IPConfigurations: &[]network.InterfaceIPConfiguration{
				{
					Name: to.StringPtr("ipConfig"),
					InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
						Subnet:                          subnet,
						PrivateIPAllocationMethod:       network.Dynamic,
						PublicIPAddress:                 ip,
						LoadBalancerBackendAddressPools: pools,
					},
				},
			},
		},
	}

	nsg, err := a.GetInstanceSecurityGroup(ctx)
	if err != nil {
		return
	}

	nicParams.NetworkSecurityGroup = &nsg

	nicClient, err := a.getNicClient()
	if err != nil {
		return
	}

	future, err := nicClient.CreateOrUpdate(ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		a.Input.InstanceInfo.NIC, nicParams)
	if err != nil {
		return
	}

	err = future.WaitForCompletionRef(ctx, nicClient.Client)
	if err != nil {
		return
	}

	return future.Result(*nicClient)
}
