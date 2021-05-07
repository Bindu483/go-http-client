package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (a *Provider) getTags(tags AWSTags) AWSTags {
	defaultTags := AWSTags{}
	defaultTags["created_by"] = "reynen court"
	defaultTags["consumer_id"] = a.ConsumerKey
	defaultTags["deployment_name"] = a.DeploymentName
	for k, v := range tags {
		defaultTags[k] = v
	}
	return defaultTags
}

func addEc2Tags(tag AWSTags, ec2Srv *ec2.EC2, resourceIds []*string) error {
	_, err := ec2Srv.CreateTags(&ec2.CreateTagsInput{
		Tags:      tag.ConvertToEc2Tags(),
		Resources: resourceIds,
	})
	return err
}
