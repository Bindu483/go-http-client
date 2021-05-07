package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
)

func New(region, deployment, consumerKey string) (*Provider, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		logrus.Error("failed to create aws session ", err.Error())
		return nil, err
	}

	_, err = s.Config.Credentials.Get()
	if err != nil {
		logrus.Error("aws credentials are not set ", err.Error())
		return nil, err
	}

	return &Provider{
		session:        s,
		DeploymentName: deployment,
		ConsumerKey:    consumerKey,
		Region:         region,
	}, nil
}
