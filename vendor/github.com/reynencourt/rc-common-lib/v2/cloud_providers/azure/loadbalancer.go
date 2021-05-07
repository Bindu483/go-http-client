package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/reynencourt/rc-common-lib/v2/ops/utils"
)

func (a *AzureProvider) getLBClient() (*network.LoadBalancersClient, error) {
	lbClient := network.NewLoadBalancersClient(a.Input.ConnectionInfo.SubscriptionID)

	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}
	lbClient.Authorizer = auth
	return &lbClient, nil
}

// GetLoadBalancer gets info on a loadbalancer
func (a *AzureProvider) GetLoadBalancer(ctx context.Context) (*network.LoadBalancer, error) {
	lbClient, err := a.getLBClient()
	if err != nil {
		return nil, err
	}

	lb, err := lbClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.LoadBalancer.Name, "")
	if err != nil {
		return nil, err
	}

	return &lb, nil
}

// CreateLoadBalancer creates a load balancer with 2 inbound NAT rules.
func (a *AzureProvider) CreateLoadBalancer(ctx context.Context) (lb network.LoadBalancer, err error) {

	probeName := fmt.Sprintf("probe-%v", utils.RandStringRunes(6))
	frontEndIPConfigName := fmt.Sprintf("fip-%v", utils.RandStringRunes(6))
	backEndAddressPoolName := fmt.Sprintf("backendPool")
	idPrefix := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/loadBalancers",
		a.Input.ConnectionInfo.SubscriptionID,
		a.Input.ConnectionInfo.ResourceGroupName)

	pip, err := a.CreatePublicIP(ctx, a.Input.LoadBalancer.PIP)
	if err != nil {
		return
	}

	lbClient, err := a.getLBClient()
	if err != nil {
		return
	}

	name := a.Input.LoadBalancer.Name
	future, err := lbClient.CreateOrUpdate(ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		name,
		network.LoadBalancer{
			Location: to.StringPtr(a.Input.ConnectionInfo.Location),
			LoadBalancerPropertiesFormat: &network.LoadBalancerPropertiesFormat{
				FrontendIPConfigurations: &[]network.FrontendIPConfiguration{
					{
						Name: &frontEndIPConfigName,
						FrontendIPConfigurationPropertiesFormat: &network.FrontendIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: network.Dynamic,
							PublicIPAddress:           pip,
						},
					},
				},
				BackendAddressPools: &[]network.BackendAddressPool{
					{
						Name: &backEndAddressPoolName,
					},
				},
				Probes: &[]network.Probe{
					{
						Name: &probeName,
						ProbePropertiesFormat: &network.ProbePropertiesFormat{
							Protocol:          network.ProbeProtocolHTTP,
							Port:              to.Int32Ptr(80),
							IntervalInSeconds: to.Int32Ptr(15),
							NumberOfProbes:    to.Int32Ptr(4),
							RequestPath:       to.StringPtr("/"),
						},
					},
				},
				LoadBalancingRules: &[]network.LoadBalancingRule{
					{
						Name: to.StringPtr("lbRule"),
						LoadBalancingRulePropertiesFormat: &network.LoadBalancingRulePropertiesFormat{
							Protocol:             network.TransportProtocolTCP,
							FrontendPort:         to.Int32Ptr(80),
							BackendPort:          to.Int32Ptr(80),
							IdleTimeoutInMinutes: to.Int32Ptr(4),
							EnableFloatingIP:     to.BoolPtr(false),
							LoadDistribution:     network.Default,
							FrontendIPConfiguration: &network.SubResource{
								ID: to.StringPtr(fmt.Sprintf("/%s/%s/frontendIPConfigurations/%s", idPrefix, name, frontEndIPConfigName)),
							},
							BackendAddressPool: &network.SubResource{
								ID: to.StringPtr(fmt.Sprintf("/%s/%s/backendAddressPools/%s", idPrefix, name, backEndAddressPoolName)),
							},
							Probe: &network.SubResource{
								ID: to.StringPtr(fmt.Sprintf("/%s/%s/probes/%s", idPrefix, name, probeName)),
							},
						},
					},
				},
			},
		})

	if err != nil {
		return lb, fmt.Errorf("cannot create load balancer: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, lbClient.Client)
	if err != nil {
		return lb, fmt.Errorf("cannot get load balancer create or update future response: %v", err)
	}

	lb, err = future.Result(*lbClient)
	if err != nil {
		return lb, err
	}

	return lb, nil
}
