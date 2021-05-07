package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func (a *AzureProvider) getAvailabilitySetsClient() (*compute.AvailabilitySetsClient, error) {
	asClient := compute.NewAvailabilitySetsClient(a.Input.ConnectionInfo.SubscriptionID)

	token, err := a.getToken()
	if err != nil {
		return nil, err
	}

	asClient.Authorizer = token

	return &asClient, nil
}

// CreateAvailabilitySet creates an availability set
func (a *AzureProvider) CreateAvailabilitySet(ctx context.Context) (*compute.AvailabilitySet, error) {
	asClient, err := a.getAvailabilitySetsClient()
	if err != nil {
		return nil, err
	}

	av, err := asClient.CreateOrUpdate(ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		a.Input.InstanceInfo.AvailabilityGroup,
		compute.AvailabilitySet{
			Location: to.StringPtr(a.Input.ConnectionInfo.Location),
			AvailabilitySetProperties: &compute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(1),
				PlatformUpdateDomainCount: to.Int32Ptr(1),
			},
			Sku: &compute.Sku{
				Name: to.StringPtr("Aligned"),
			},
		})

	if err != nil {
		return nil, err
	}

	return &av, nil
}

// GetAvailabilitySet gets info on an availability set
func (a *AzureProvider) GetAvailabilitySet(ctx context.Context) (*compute.AvailabilitySet, error) {
	asClient, err := a.getAvailabilitySetsClient()
	if err != nil {
		return nil, err
	}

	av, err := asClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.InstanceInfo.AvailabilityGroup)
	if err != nil {
		return nil, err
	}

	return &av, nil
}
