package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

func (a *AzureProvider) createGroupClient() (*resources.GroupsClient, error) {
	groupsClient := resources.NewGroupsClient(a.Input.ConnectionInfo.SubscriptionID)

	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}

	groupsClient.Authorizer = auth

	return &groupsClient, nil
}

func (a *AzureProvider) GetResourceGroup(ctx context.Context) (*resources.Group, error) {
	rg := resources.Group{
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
	}

	groupClient, err := a.createGroupClient()
	if err != nil {
		return nil, err
	}

	rg, err = groupClient.Get(ctx, a.Input.ConnectionInfo.ResourceGroupName)
	if err != nil {
		return nil, err
	}

	return &rg, nil
}

func (a *AzureProvider) CreateResourceGroup(ctx context.Context) (*resources.Group, error) {

	rg := resources.Group{
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
	}

	groupClient, err := a.createGroupClient()
	if err != nil {
		return nil, err
	}

	_, err = groupClient.CreateOrUpdate(ctx, a.Input.ConnectionInfo.ResourceGroupName, rg)
	if err != nil {
		return nil, err
	}

	return &rg, nil
}

func (a *AzureProvider) DeleteResourceGroup(ctx context.Context) error {

	groupClient, err := a.createGroupClient()
	if err != nil {
		return err
	}

	futures, err := groupClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName)
	if err != nil {
		return err
	}

	err = futures.WaitForCompletionRef(ctx, groupClient.Client)
	if err != nil {
		return err
	}

	return nil
}
