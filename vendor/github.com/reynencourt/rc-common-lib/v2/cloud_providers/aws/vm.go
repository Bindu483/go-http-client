package aws

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
)

type CreateInstanceRequest struct {
	Count                 int64
	SubnetId              string
	InstanceType          string
	OsType                cloud_providers.OsType
	KeyName               string
	Tags                  AWSTags
	SecurityGroupIds      []string
	IamInstanceProfileArn string
	DiskSize              int64
}

var (
	ErrOverProvisioning = errors.New("could not acquire requested amount of machines")
)

type CreateInstanceResponse struct {
	PrivateIp  string
	PublicIp   string
	InstanceId string
}

const (
	InvalidRegion       = "invalid region"
	InvalidInstanceType = "invalid instance type"
	InvalidOsType       = "invalid os type"
	InstanceCreateErr   = "failed to create instance"
)

func (a *Provider) CreateInstances(req CreateInstanceRequest) (Instances, error) {

	if yes := IsValidAwsRegion(a.Region); !yes {
		return nil, errors.New(InvalidRegion)
	}

	if yes := IsApprovedInstanceType(req.InstanceType); !yes {
		return nil, errors.New(InvalidInstanceType)
	}

	if yes := IsValidOsType(req.OsType); !yes {
		return nil, errors.New(InvalidOsType)
	}

	var instances Instances

	if req.Count == 0 {
		return nil, errors.New("count cant be zero")
	}

	ec2Srv := ec2.New(a.session)

	//https://forums.aws.amazon.com/thread.jspa?messageID=593651
	updatedTags := req.Tags
	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = "rc-platform"
	}

	var ec2Out *ec2.Reservation
	var ec2CreationError error

	currTime := time.Now()

	for time.Since(currTime) < 15*time.Minute {
		blockDevice := getBlockDeviceMappings(a.Region, req.OsType)
		if len(blockDevice) > 0 && req.DiskSize >= 20 {
			blockDevice[0].Ebs.VolumeSize = &req.DiskSize
		}

		ec2Out, ec2CreationError = ec2Srv.RunInstances(&ec2.RunInstancesInput{
			SubnetId:            aws.String(req.SubnetId),
			MinCount:            aws.Int64(1),
			ImageId:             aws.String(getAmiId(a.Region, req.OsType)),
			InstanceType:        aws.String(req.InstanceType),
			MaxCount:            aws.Int64(req.Count),
			BlockDeviceMappings: blockDevice,
			KeyName:             aws.String(req.KeyName),
			SecurityGroupIds:    aws.StringSlice(req.SecurityGroupIds),
			TagSpecifications: []*ec2.TagSpecification{
				{
					Tags:         a.getTags(updatedTags).ConvertToEc2Tags(),
					ResourceType: aws.String("instance"),
				},
			},
			IamInstanceProfile: &ec2.IamInstanceProfileSpecification{Arn: aws.String(req.IamInstanceProfileArn)},
		})
		if ec2CreationError != nil {
			if awsErr, ok := ec2CreationError.(awserr.Error); ok {
				if awsErr.Code() == "InvalidParameterValue" {
					time.Sleep(5 * time.Second)
					continue
				} else {
					break
				}
			} else {
				break
			}
		}
		ec2CreationError = nil
		break
	}

	if ec2CreationError != nil {
		return nil, ec2CreationError
	}

	if ec2Out == nil {
		return nil, errors.New("unknown error happened")
	}

	var waitForEc2Instances sync.WaitGroup
	var err error

	waitForEc2Instances.Add(int(req.Count))

	if len(ec2Out.Instances) != int(req.Count) {

		var instanceIDs []string
		for _, inst := range ec2Out.Instances {
			instanceIDs = append(instanceIDs, aws.StringValue(inst.InstanceId))
		}

		//Best effort in cleaning up
		_ = a.DeleteInstances(instanceIDs)

		return nil, ErrOverProvisioning
	}

	for _, instance := range ec2Out.Instances {

		go func(awsInstance *ec2.Instance) {

			for {

				if time.Since(currTime) > 15*time.Minute {

					for i := 0; i < len(ec2Out.Instances); i++ {
						waitForEc2Instances.Done()
					}

					err = errors.New("could not acquire VM")
					return
				}

				if awsInstance == nil {
					waitForEc2Instances.Done()
					return
				}

				resp, err := ec2Srv.DescribeInstances(&ec2.DescribeInstancesInput{
					InstanceIds: []*string{awsInstance.InstanceId},
				})

				if err != nil {
					continue
				}

				if len(resp.Reservations) < 1 {
					continue
				}
				if len(resp.Reservations[0].Instances) < 1 {
					continue
				}
				if resp.Reservations[0].Instances[0].State == nil {
					continue
				}
				if resp.Reservations[0].Instances[0].State.Name == nil {
					continue
				}

				state := *resp.Reservations[0].Instances[0].State.Name

				switch state {
				case ec2.InstanceStateNameRunning:
					// We're ready!

					var name string
					tags := resp.Reservations[0].Instances[0].Tags

					for _, t := range tags {
						if aws.StringValue(t.Key) == "Name" {
							name = aws.StringValue(t.Value)
						}
					}

					if awsInstance != nil {
						instances = append(instances, &Instance{
							Name:       name,
							InstanceId: aws.StringValue(awsInstance.InstanceId),
							PublicIp:   aws.StringValue(resp.Reservations[0].Instances[0].PublicIpAddress),
							PrivateIp:  aws.StringValue(resp.Reservations[0].Instances[0].PrivateIpAddress),
						})
					}

					waitForEc2Instances.Done()
					return
				case ec2.InstanceStateNameTerminated, ec2.InstanceStateNameStopped,
					ec2.InstanceStateNameStopping, ec2.InstanceStateNameShuttingDown:
					waitForEc2Instances.Done()
					return
				}

				time.Sleep(2 * time.Second)
			}

		}(instance)
	}

	waitForEc2Instances.Wait()

	if err != nil {
		return instances, err
	}

	return instances, nil
}

func (a *Provider) DeleteInstances(ids []string) error {
	ec2Srv := ec2.New(a.session)

	instances, err := ec2Srv.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: aws.StringSlice(ids),
	})
	if err != nil {
		return err
	}
	var waitForEc2Instances sync.WaitGroup

	waitForEc2Instances.Add(len(ids))

	for _, instance := range instances.TerminatingInstances {

		go func(awsInstance *ec2.InstanceStateChange) {

			for {

				if awsInstance == nil {
					waitForEc2Instances.Done()
					return
				}

				resp, err := ec2Srv.DescribeInstances(&ec2.DescribeInstancesInput{
					InstanceIds: []*string{awsInstance.InstanceId},
				})

				if err != nil {
					fmt.Printf("%s", err.Error())
					waitForEc2Instances.Done()
					return
				}

				state := *resp.Reservations[0].Instances[0].State.Name

				switch state {
				case ec2.InstanceStateNameTerminated:
					waitForEc2Instances.Done()
					return
				}

				time.Sleep(2 * time.Second)
			}

		}(instance)
	}

	waitForEc2Instances.Wait()
	return nil
}

func (a *Provider) DeleteInstancesByIp(ips []string, vpcId string) error {
	ec2Srv := ec2.New(a.session)

	out, err := ec2Srv.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{Name: aws.String("vpc-id"), Values: aws.StringSlice([]string{vpcId})},
			{Name: aws.String("private-ip-address"), Values: aws.StringSlice(ips)},
		}, MaxResults: aws.Int64(1000),
	})
	if err != nil {
		return err
	}
	var instanceIds []string

	for _, r := range out.Reservations {
		for _, i := range r.Instances {
			instanceIds = append(instanceIds, *i.InstanceId)
		}
	}

	return a.DeleteInstances(instanceIds)
}

func GetInstanceTypes() []InstanceType {
	return approvedInstanceTypes
}

func GetOsTypes() []cloud_providers.OsType { return approvedOsTypes }

func GetRegions() []string { return awsRegions }

type InstanceType struct {
	Type           string
	NoCpu          int
	MemoryCapacity int
}
