package aws

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/iam"
	"strconv"
	"strings"
	"time"
)

type ClassicLbRequest struct {
	Name              string
	SubnetId          []string
	SecurityGroupIds  []string
	InstancesToAttach []string
	InstancePort      int64
	LbPort            int64
	HealthCheckPath   string
	Scheme            LBScheme
	Tags              AWSTags
	Protocol          Protocol
}
type ApplicationLoadBalancerCreationRequest struct {
	Name              string
	IpAddressType     IpAddressType //defaults to ipv4
	SubnetId          []string
	SecurityGroupIds  []string
	InstancesToAttach []string
	InstancePorts     []uint32
	HealthCheckPath   string
	Tags              AWSTags
	TLSMetadata       []TLSMetadata
	Scheme            LBScheme
	VpcId             string
}

type TLSMetadata struct {
	Cert string
	Key  string
	Name string
}

type IpAddressType string

const (
	IpAddressTypeIpv4      IpAddressType = "ipv4"
	IpAddressTypeDualStack IpAddressType = "dualstack"
)

type Protocol string

const (
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
	TCP   Protocol = "TCP"
)

type LBScheme string

const (
	LBSchemePublic  LBScheme = "internet-facing"
	LBSchemePrivate LBScheme = "internal"
)

type Err func(err error) error

var errLBAttachInstanceErr Err = func(err error) error {
	return fmt.Errorf("failed to attach instances to lb - %v", err)
}

const (
	LbCreateErr            = "failed to create load balancer"
	LbCreateHealthCheckErr = "failed to create lb health check"
	LbDescribeErr          = "failed to get lb names"
)

func (a *Provider) CreateApplicationLoadBalancer(req *ApplicationLoadBalancerCreationRequest) (string, error) {
	client := elbv2.New(a.session)
	tags := req.Tags
	if len(tags) == 0 {
		tags = AWSTags{}
	}
	if tags["Name"] == "" {
		delete(tags, "name")
		tags["Name"] = req.Name
	}
	var createdResources []string
	lbName := req.Name
	if len(lbName) > 32 {
		lbName = lbName[len(lbName)-31:]
	}
	lBCreationRequest := elbv2.CreateLoadBalancerInput{
		IpAddressType:  aws.String(string(req.IpAddressType)),
		Name:           aws.String(lbName),
		Subnets:        aws.StringSlice(req.SubnetId),
		Scheme:         aws.String(string(req.Scheme)),
		SecurityGroups: aws.StringSlice(req.SecurityGroupIds),
		Tags:           tags.ConvertToElbV2Tags(),
		Type:           aws.String(elbv2.LoadBalancerTypeEnumApplication),
	}
	loadBalancerCreationOutput, err := client.CreateLoadBalancer(&lBCreationRequest)
	if err != nil {
		return "", errors.New(LbCreateErr + err.Error())
	}
	if len(loadBalancerCreationOutput.LoadBalancers) == 0 {
		return "", errors.New(LbCreateErr + " aws output is empty")
	}
	loadBalancer := loadBalancerCreationOutput.LoadBalancers[0] //we are only creating one lb

	createdResources = append(createdResources, *loadBalancer.LoadBalancerArn)

	// create a tg for every unique port and add instances to it
	for _, port := range req.InstancePorts {
		tgName := fmt.Sprintf("%s-tg-%d", lbName, port)
		if len(tgName) > 32 {
			tgName = tgName[len(tgName)-31:]
		}
		var protocolToBeUsed string
		if port == 80 {
			protocolToBeUsed = elbv2.ProtocolEnumHttp
		} else if port == 443 {
			protocolToBeUsed = elbv2.ProtocolEnumHttps
		} else {
			return "", errors.New("unsupported protocol for creating the target group")
		}
		tgCreationOutput, err := client.CreateTargetGroup(&elbv2.CreateTargetGroupInput{
			HealthCheckPath:     aws.String(req.HealthCheckPath),
			HealthCheckPort:     aws.String(strconv.Itoa(int(port))),
			HealthCheckProtocol: aws.String(elbv2.ProtocolEnumHttp),
			Name:                aws.String(tgName),
			Port:                aws.Int64(int64(port)),
			Protocol:            aws.String(protocolToBeUsed),
			VpcId:               aws.String(req.VpcId),
		})
		if err != nil {
			return "", errors.New(LbCreateErr + err.Error())
		}
		if len(tgCreationOutput.TargetGroups) == 0 {
			return "", errors.New(LbCreateErr + " unable to create target groups")
		}
		tg := tgCreationOutput.TargetGroups[0]

		// just in case we have more tgs created here
		for _, t := range tgCreationOutput.TargetGroups {
			createdResources = append(createdResources, *t.TargetGroupArn)
		}
		var targetDescriptions []*elbv2.TargetDescription
		for _, ec2InstanceId := range req.InstancesToAttach {
			targetDescriptions = append(targetDescriptions, &elbv2.TargetDescription{
				Id:   aws.String(ec2InstanceId),
				Port: aws.Int64(int64(port)),
			})
		}

		_, err = client.RegisterTargets(&elbv2.RegisterTargetsInput{
			TargetGroupArn: tg.TargetGroupArn,
			Targets:        targetDescriptions,
		})
		if err != nil {
			return "", errors.New(LbCreateErr + err.Error())
		}

		// create a forward action
		// we can actually forward port 80 the traffic to 443, however for the sake of cleanliness, I'm keeping it to the same port
		// + the ssl redirection is already done at the browser level with the upgrade headers.
		forwardAction := elbv2.Action{

			TargetGroupArn: tg.TargetGroupArn,
			Type:           aws.String(elbv2.ActionTypeEnumForward),
		}
		listenerCreationRequest := elbv2.CreateListenerInput{
			DefaultActions:  []*elbv2.Action{&forwardAction},
			LoadBalancerArn: loadBalancer.LoadBalancerArn,
			Port:            aws.Int64(int64(port)),
			Protocol:        aws.String(protocolToBeUsed),
			//Certificates:    nil,
		}
		if port == 443 {
			iamClient := iam.New(a.session)
			var awsCerts []*elbv2.Certificate
			for _, cert := range req.TLSMetadata {
				input := &iam.UploadServerCertificateInput{
					CertificateBody:       aws.String(cert.Cert),
					PrivateKey:            aws.String(cert.Key),
					ServerCertificateName: aws.String(cert.Name),
				}
				certificateCreationResponse, err := iamClient.UploadServerCertificate(input)
				if err != nil {
					return "", errors.New(LbCreateErr + " upload certificate to iam returned error. " + err.Error())
				}

				awsCerts = append(awsCerts, &elbv2.Certificate{
					CertificateArn: certificateCreationResponse.ServerCertificateMetadata.Arn,
				})
			}

			listenerCreationRequest.Certificates = awsCerts
			listenerCreationRequest.SslPolicy = aws.String("ELBSecurityPolicy-2016-08")
		}
		listenerCreationRequest.SetProtocol(protocolToBeUsed)
		retryCount := 10
		valid := false
		for i := retryCount; i > 0; i-- {
			_, err := client.CreateListener(&listenerCreationRequest)
			if err == nil {
				valid = true
				break
			} else {
				if isAWSErr(err, elbv2.ErrCodeCertificateNotFoundException, "") {
					fmt.Println("going to retry after 30 seconds...")
					time.Sleep(time.Second * 30)
				} else {
					// should not continue if it is not 404
					fmt.Println(err)
					break
				}

			}
		}
		if !valid {
			return "", errors.New(LbCreateErr + "unable to create the listener")
		}

	}

	return *loadBalancer.DNSName, nil

}

// Returns true if the error matches all these conditions:
//  * err is of type awserr.Error
//  * Error.Code() matches code
//  * Error.Message() contains message
func isAWSErr(err error, code string, message string) bool {
	var awsErr awserr.Error
	if errors.As(err, &awsErr) {
		return awsErr.Code() == code && strings.Contains(awsErr.Message(), message)
	}
	return false
}

// TODO: Rollback function
//func (a *Provider) rollback(client *elbv2.ELBV2, resources []string) error {
//
//}
func (a *Provider) CreateClassicLoadBalancer(req *ClassicLbRequest) (lbDnsName string, err error) {
	elbSrv := elb.New(a.session)
	updatedTags := req.Tags
	if len(updatedTags) == 0 {
		updatedTags = AWSTags{}
	}
	if updatedTags["Name"] == "" {
		delete(updatedTags, "name")
		updatedTags["Name"] = req.Name
	}

	lbName := req.Name
	if len(lbName) > 32 {
		lbName = lbName[0:31]
	}
	out, err := elbSrv.CreateLoadBalancer(&elb.CreateLoadBalancerInput{
		LoadBalancerName: aws.String(lbName),
		Tags:             updatedTags.ConvertToElbTags(),
		Listeners: []*elb.Listener{
			{
				Protocol:         aws.String("tcp"),
				InstancePort:     aws.Int64(req.InstancePort),
				InstanceProtocol: aws.String("tcp"),
				LoadBalancerPort: aws.Int64(req.LbPort),
			},
		},
		SecurityGroups: aws.StringSlice(req.SecurityGroupIds),
		Subnets:        aws.StringSlice(req.SubnetId),
		Scheme:         aws.String(string(req.Scheme)),
	})
	if err != nil {
		return "", errors.New(LbCreateErr + err.Error())
	}

	var ec2Instances []*elb.Instance

	for _, i := range req.InstancesToAttach {
		ec2Instances = append(ec2Instances, &elb.Instance{InstanceId: aws.String(i)})
	}

	_, err = elbSrv.RegisterInstancesWithLoadBalancer(&elb.RegisterInstancesWithLoadBalancerInput{
		LoadBalancerName: aws.String(req.Name),
		Instances:        ec2Instances,
	})
	if err != nil {
		return *out.DNSName, errLBAttachInstanceErr(err)
	}

	_, err = elbSrv.ConfigureHealthCheck(&elb.ConfigureHealthCheckInput{
		LoadBalancerName: aws.String(req.Name),
		HealthCheck: &elb.HealthCheck{
			HealthyThreshold:   aws.Int64(2),
			UnhealthyThreshold: aws.Int64(2),
			Interval:           aws.Int64(6),
			Target:             aws.String(fmt.Sprintf("%s:%v/%s", req.Protocol, req.InstancePort, req.HealthCheckPath)),
			Timeout:            aws.Int64(5),
		},
	})

	if err != nil {
		return *out.DNSName, errors.New(LbCreateHealthCheckErr + err.Error())
	}

	return *out.DNSName, nil
}

func (a *Provider) DoesClassicLbExist(lbName string) (bool, error) {
	elbSrv := elb.New(a.session)
	out, err := elbSrv.DescribeLoadBalancers(&elb.DescribeLoadBalancersInput{LoadBalancerNames: aws.StringSlice([]string{lbName}), PageSize: aws.Int64(100)})
	if err != nil {
		return false, errors.New(LbDescribeErr + err.Error())
	}
	for _, j := range out.LoadBalancerDescriptions {
		if *j.LoadBalancerName == lbName {
			return true, nil
		}
	}
	return false, nil
}
