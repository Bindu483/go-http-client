package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
)

type AWSTags map[string]string

func (t AWSTags) ConvertToEc2Tags() []*ec2.Tag {
	var tags []*ec2.Tag
	for k, v := range t {
		tags = append(tags, &ec2.Tag{Key: aws.String(k), Value: aws.String(v)})
	}
	return tags
}

func (t AWSTags) ConvertToElbTags() []*elb.Tag {
	var tags []*elb.Tag

	for key, value := range t {
		tags = append(tags, &elb.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return tags
}
func (t AWSTags) ConvertToElbV2Tags() []*elbv2.Tag {
	var tags []*elbv2.Tag

	for key, value := range t {
		tags = append(tags, &elbv2.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return tags
}

func (t AWSTags) ConvertToIamTags() []*iam.Tag {
	var tags []*iam.Tag

	for key, value := range t {
		tags = append(tags, &iam.Tag{Key: aws.String(key), Value: aws.String(value)})
	}
	return tags
}

var awsRegions = []string{
	"eu-north-1",
	"ap-south-1",
	"eu-west-3",
	"eu-west-2",
	"eu-west-1",
	"ap-northeast-2",
	"ap-northeast-1",
	"sa-east-1",
	"ca-central-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"eu-central-1",
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

var approvedInstanceTypes = []InstanceType{
	{
		Type:           "t2.xlarge",
		MemoryCapacity: 16,
		NoCpu:          4,
	},
	{
		Type:           "t2.2xlarge",
		MemoryCapacity: 32,
		NoCpu:          8,
	},
	{
		Type:           "t2.large",
		MemoryCapacity: 8,
		NoCpu:          2,
	},
	{
		Type:           "t3.large",
		MemoryCapacity: 8,
		NoCpu:          2,
	}, {
		Type:           "t3.xlarge",
		MemoryCapacity: 16,
		NoCpu:          4,
	}, {
		Type:           "t3.2xlarge",
		MemoryCapacity: 32,
		NoCpu:          8,
	}, {
		Type:           "t3a.large",
		MemoryCapacity: 8,
		NoCpu:          2,
	}, {
		Type:           "t3a.xlarge",
		MemoryCapacity: 16,
		NoCpu:          4,
	}, {
		Type:           "t3a.2xlarge",
		MemoryCapacity: 32,
		NoCpu:          8,
	},
}

var approvedOsTypes = []cloud_providers.OsType{
	cloud_providers.Ubuntu1804,
}

var AMIs = map[string]map[cloud_providers.OsType]string{
	"us-east-2": {
		cloud_providers.Ubuntu1804: "ami-0f65671a86f061fcd",
	},
	"us-east-1": {
		cloud_providers.Ubuntu1804: "ami-0817d428a6fb68645",
	},
	"us-west-1": {
		cloud_providers.Ubuntu1804: "ami-03fac5402e10ea93b",
	},
	"us-west-2": {
		cloud_providers.Ubuntu1804: "ami-07a29e5e945228fa1",
	},
	"eu-central-1": {
		cloud_providers.Ubuntu1804: "ami-092391a11f8aa4b7b",
	},
	"eu-west-1": {
		cloud_providers.Ubuntu1804: "ami-0823c236601fef765",
	},
	"eu-west-2": {
		cloud_providers.Ubuntu1804: "ami-09a1e275e350acf38",
	},
	"eu-west-3": {
		cloud_providers.Ubuntu1804: "ami-014d8dccd70fd2632",
	},
	"eu-north-1": {
		cloud_providers.Ubuntu1804: "ami-0ede7f804d699ea83",
	},
}

var blockDeviceMapping = map[string]map[cloud_providers.OsType]*ec2.BlockDeviceMapping{
	"us-east-2": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"us-east-1": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"us-west-1": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"us-west-2": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"eu-central-1": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"eu-west-1": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"eu-west-2": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"eu-west-3": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
	"eu-north-1": {
		cloud_providers.Ubuntu1804: {
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize:          aws.Int64(100),
				DeleteOnTermination: aws.Bool(true),
				VolumeType:          aws.String("gp2"),
			},
		},
	},
}

func IsValidAwsRegion(r string) bool {
	for _, region := range awsRegions {
		if region == r {
			return true
		}
	}
	return false
}

func IsValidOsType(o cloud_providers.OsType) bool {
	for _, oType := range approvedOsTypes {
		if oType == o {
			return true
		}
	}
	return false
}

func IsApprovedInstanceType(instanceType string) bool {
	for _, iType := range approvedInstanceTypes {
		if instanceType == iType.Type {
			return true
		}
	}
	return false
}

func getAmiId(region string, osType cloud_providers.OsType) string {
	return AMIs[region][osType]
}

func getBlockDeviceMappings(region string, osType cloud_providers.OsType) []*ec2.BlockDeviceMapping {
	return []*ec2.BlockDeviceMapping{blockDeviceMapping[region][osType]}
}

func (i Instances) GetInstanceIds() []string {
	var ids []string
	for _, instance := range i {
		ids = append(ids, instance.InstanceId)
	}
	return ids
}

func (i Instances) GetInstancePrivateIps() []string {
	var ips []string
	for _, instance := range i {
		ips = append(ips, instance.PrivateIp)
	}
	return ips
}

func (i Instances) GetInstancePublicIps() []string {
	var ips []string
	for _, instance := range i {
		ips = append(ips, instance.PublicIp)
	}
	return ips
}

func (f FireWallRuleOutputs) GetFirewalId(name string) string {
	for _, sg := range f {
		if sg.Name == name {
			return sg.FirewallRuleId
		}
	}
	return ""
}

func (s Subnets) GetSubnetId(name string) string {
	for _, sub := range s {
		if sub.Name == name {
			return sub.Id
		}
	}
	return ""
}
