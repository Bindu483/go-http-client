package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func (a *Provider) CreateRcIamPolicy(policyName string) (string, error) {
	iamSrv := iam.New(a.session)
	out, err := iamSrv.CreatePolicy(&iam.CreatePolicyInput{
		Description:    aws.String("iam policy for rc user"),
		PolicyName:     aws.String(policyName),
		PolicyDocument: aws.String(RcUserIamPolicy),
	})
	if err != nil {
		return "", err
	}
	return *out.Policy.Arn, err
}

func (a *Provider) DeleteRcIamPolicy(policyArn string) error {
	iamSrv := iam.New(a.session)

	_, err := iamSrv.DeletePolicy(&iam.DeletePolicyInput{PolicyArn: aws.String(policyArn)})
	return err
}

func (a *Provider) CreateRcIamRole(roleName string, policyArn string, tags AWSTags) (string, error) {
	iamSrv := iam.New(a.session)

	iamRole, err := iamSrv.CreateRole(&iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(RcAssumeRolePolicy),
		Description:              aws.String("iam role for rc platform"),
		RoleName:                 aws.String(roleName),

		PermissionsBoundary: aws.String(policyArn),
		Tags:                tags.ConvertToIamTags(),
		MaxSessionDuration:  aws.Int64(3600),
	})
	if err != nil {
		return "", err
	}
	_, err = iamSrv.AttachRolePolicy(&iam.AttachRolePolicyInput{RoleName: aws.String(roleName), PolicyArn: aws.String(policyArn)})
	if err != nil {
		return "", err
	}
	return *iamRole.Role.Arn, nil
}

func (a *Provider) DeleteRcIamRole(roleName string, rcPolicyArn string) error {
	iamSrv := iam.New(a.session)
	_, err := iamSrv.DetachRolePolicy(&iam.DetachRolePolicyInput{RoleName: aws.String(roleName), PolicyArn: aws.String(rcPolicyArn)})
	_, err = iamSrv.DeleteRole(&iam.DeleteRoleInput{RoleName: aws.String(roleName)})
	if err != nil {
		return err
	}
	return nil
}

func (a *Provider) CreateRcInstanceProfile(roleName string) (string, error) {

	var instanceProfileName = fmt.Sprintf("%s-profile", roleName)

	iamSrv := iam.New(a.session)

	profile, err := iamSrv.CreateInstanceProfile(&iam.CreateInstanceProfileInput{
		InstanceProfileName: aws.String(instanceProfileName),
	})
	if err != nil {
		return "", err
	}
	_, err = iamSrv.AddRoleToInstanceProfile(&iam.AddRoleToInstanceProfileInput{InstanceProfileName: aws.String(instanceProfileName), RoleName: aws.String(roleName)})

	if err != nil {
		return "", nil
	}

	return *profile.InstanceProfile.Arn, nil
}

func (a *Provider) DeleteRcInstanceProfile(roleName string) error {

	var instanceProfileName = fmt.Sprintf("%s-profile", roleName)

	iamSrv := iam.New(a.session)

	_, err := iamSrv.RemoveRoleFromInstanceProfile(&iam.RemoveRoleFromInstanceProfileInput{InstanceProfileName: aws.String(instanceProfileName), RoleName: aws.String(roleName)})
	if err != nil {
		return err
	}
	_, err = iamSrv.DeleteInstanceProfile(&iam.DeleteInstanceProfileInput{InstanceProfileName: aws.String(instanceProfileName)})
	if err != nil {
		return err
	}
	return nil
}
