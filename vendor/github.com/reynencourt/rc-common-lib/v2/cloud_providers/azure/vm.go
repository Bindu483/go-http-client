package azure

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/compute/mgmt/compute"
	"github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"strings"
	"sync"
)

type InstanceType struct {
	Type           string
	MemoryCapacity int
	NoCpu          int
}

var approvedInstanceTypes = []InstanceType{
	{
		Type:           "StandardDS1V2",
		MemoryCapacity: 3,
		NoCpu:          1,
	},
	{
		Type:           "StandardDS2V2",
		MemoryCapacity: 7,
		NoCpu:          2,
	},
	{
		Type:           "StandardDS3V2",
		MemoryCapacity: 14,
		NoCpu:          4,
	},
	{
		Type:           "StandardDS4V2",
		MemoryCapacity: 28,
		NoCpu:          8,
	},
	{
		Type:           "StandardDS5V2",
		MemoryCapacity: 56,
		NoCpu:          16,
	},
}

func IsApprovedInstanceType(instanceType string) bool {
	for _, iType := range approvedInstanceTypes {
		if instanceType == iType.Type {
			return true
		}
	}
	return false
}

func (a *AzureProvider) getVMClient() (*compute.VirtualMachinesClient, error) {
	vmClient := compute.NewVirtualMachinesClient(a.Input.ConnectionInfo.SubscriptionID)
	auth, err := a.getToken()
	if err != nil {
		return nil, err
	}
	vmClient.Authorizer = auth
	return &vmClient, nil
}

func (a *AzureProvider) validateInstanceInfo(instanceInfo *InstanceInfo) error {
	if instanceInfo == nil {
		return errors.New("instance info is not set")
	}

	if strings.TrimSpace(instanceInfo.Username) == "" {
		return errors.New("user name is not set")
	}

	if strings.TrimSpace(instanceInfo.Version) == "" {
		return errors.New("os version is not set")
	}

	if strings.TrimSpace(instanceInfo.Sku) == "" {
		return errors.New("os sku is not set")
	}

	if strings.TrimSpace(instanceInfo.Offer) == "" {
		return errors.New("os offer is not set")
	}

	if strings.TrimSpace(instanceInfo.Publisher) == "" {
		return errors.New("os publisher is not set")
	}

	if strings.TrimSpace(instanceInfo.SSHPublicKey) == "" {
		return errors.New("public key is not set")
	}

	if strings.TrimSpace(instanceInfo.InstanceName) == "" {
		return errors.New("instance name is not set")
	}

	return nil
}

func (a *AzureProvider) DeleteVMs(ctx context.Context, ids []string) error {

	vmClient, err := a.getVMClient()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, id := range ids {

		wg.Add(1)

		go func() {

			defer wg.Done()

			future, err := vmClient.Delete(ctx, a.Input.ConnectionInfo.ResourceGroupName, id)
			if err != nil {
				return
			}

			err = future.WaitForCompletionRef(ctx, vmClient.Client)
			if err != nil {
				return
			}
		}()
	}

	wg.Wait()

	return nil

}

func (a *AzureProvider) CreateClusterVM(ctx context.Context, instanceName string) (*compute.VirtualMachine, error) {

	var (
		err          error
		availability *compute.AvailabilitySet
	)

	vmClient, err := a.getVMClient()
	if err != nil {
		return nil, err
	}

	availability, err = a.GetAvailabilitySet(ctx)
	if err != nil {
		return nil, err
	}

	nic, err := a.CreateClusterNic(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	machine := compute.VirtualMachine{
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.StandardDS1V2,
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &compute.ImageReference{
					Publisher: to.StringPtr(a.Input.InstanceInfo.Publisher),
					Offer:     to.StringPtr(a.Input.InstanceInfo.Offer),
					Sku:       to.StringPtr(a.Input.InstanceInfo.Sku),
					Version:   to.StringPtr(a.Input.InstanceInfo.Version),
				},
			},
			OsProfile: &compute.OSProfile{
				ComputerName:  to.StringPtr(instanceName),
				AdminUsername: to.StringPtr(a.Input.InstanceInfo.Username),
				LinuxConfiguration: &compute.LinuxConfiguration{
					SSH: &compute.SSHConfiguration{
						PublicKeys: &[]compute.SSHPublicKey{
							{
								Path: to.StringPtr(
									fmt.Sprintf("/home/%s/.ssh/authorized_keys",
										a.Input.InstanceInfo.Username)),
								KeyData: to.StringPtr(a.Input.InstanceInfo.SSHPublicKey),
							},
						},
					},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: nic.ID,
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							Primary: to.BoolPtr(true),
						},
					},
				},
			},
			AvailabilitySet: &compute.SubResource{
				ID: availability.ID,
			},
		},
	}

	future, err := vmClient.CreateOrUpdate(
		ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		instanceName,
		machine,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the vm create or update future response: %v", err)
	}

	vm, err := future.Result(*vmClient)
	if err != nil {
		return nil, err
	}

	return &vm, nil

}

func (a *AzureProvider) CreateRcVM(ctx context.Context) (*compute.VirtualMachine, error) {

	var err error
	var availability *compute.AvailabilitySet

	if err := a.validateInstanceInfo(a.Input.InstanceInfo); err != nil {
		return nil, err
	}

	_, err = a.CreatePublicIP(ctx, a.Input.InstanceInfo.IPName)
	if err != nil {
		return nil, err
	}

	if a.Input.InstanceInfo.AvailabilityGroup != "" {
		availability, err = a.GetAvailabilitySet(ctx)
		if err != nil {
			return nil, err
		}
	}

	_, err = a.CreateInstanceSecurityGroup(ctx)
	if err != nil {
		return nil, err
	}

	var nic network.Interface

	nic, err = a.CreateInstanceNIC(ctx, a.Input.InstanceInfo.IPName)
	if err != nil {
		return nil, err
	}

	var sshKeyData = a.Input.InstanceInfo.SSHPublicKey

	vmClient, err := a.getVMClient()
	if err != nil {
		return nil, err
	}

	machine := compute.VirtualMachine{
		Location: to.StringPtr(a.Input.ConnectionInfo.Location),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.StandardDS1V2,
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &compute.ImageReference{
					Publisher: to.StringPtr(a.Input.InstanceInfo.Publisher),
					Offer:     to.StringPtr(a.Input.InstanceInfo.Offer),
					Sku:       to.StringPtr(a.Input.InstanceInfo.Sku),
					Version:   to.StringPtr(a.Input.InstanceInfo.Version),
				},
			},
			OsProfile: &compute.OSProfile{
				ComputerName:  to.StringPtr(a.Input.InstanceInfo.InstanceName),
				AdminUsername: to.StringPtr(a.Input.InstanceInfo.Username),
				LinuxConfiguration: &compute.LinuxConfiguration{
					SSH: &compute.SSHConfiguration{
						PublicKeys: &[]compute.SSHPublicKey{
							{
								Path: to.StringPtr(
									fmt.Sprintf("/home/%s/.ssh/authorized_keys",
										a.Input.InstanceInfo.Username)),
								KeyData: to.StringPtr(sshKeyData),
							},
						},
					},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: nic.ID,
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							Primary: to.BoolPtr(true),
						},
					},
				},
			},
		},
	}

	if availability != nil {
		machine.AvailabilitySet = &compute.SubResource{
			ID: availability.ID,
		}
	}

	future, err := vmClient.CreateOrUpdate(
		ctx,
		a.Input.ConnectionInfo.ResourceGroupName,
		a.Input.InstanceInfo.InstanceName,
		machine,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the vm create or update future response: %v", err)
	}

	vm, err := future.Result(*vmClient)
	if err != nil {
		return nil, err
	}

	return &vm, nil
}

// DeallocateVM deallocates the selected VM
func (a *AzureProvider) DeallocateVM(ctx context.Context) (*compute.OperationStatusResponse, error) {
	vmClient, err := a.getVMClient()
	if err != nil {
		return nil, err
	}

	future, err := vmClient.Deallocate(ctx, a.Input.ConnectionInfo.ResourceGroupName, a.Input.InstanceInfo.InstanceName)
	if err != nil {
		return nil, fmt.Errorf("cannot deallocate vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return nil, fmt.Errorf("cannot get the vm deallocate future response: %v", err)
	}

	resp, err := future.Result(*vmClient)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
