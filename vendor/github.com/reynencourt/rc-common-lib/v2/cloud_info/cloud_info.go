package cloud_info

import (
	"encoding/json"
	"errors"

	"github.com/reynencourt/rc-common-lib/v2/cloud_providers"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers/aws"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers/azure"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers/onprem"
)

func MarshalCloudInfo(info *cloud_providers.RcInfo) ([]byte, error) {
	return json.Marshal(info)
}

func GetSMTPCredential(b []byte) (*cloud_providers.SMTPCredential, error) {
	cloudInfo, err := UnMarshalCloudInfo(b)
	if err != nil {
		return nil, err
	}
	return &cloudInfo.SMTPCredential, nil
}

func SetSMTPCredentials(b []byte, s cloud_providers.SMTPCredential) ([]byte, error) {
	cloudInfo, err := UnMarshalCloudInfo(b)
	if err != nil {
		return nil, err
	}
	cloudInfo.SMTPCredential = s
	return MarshalCloudInfo(cloudInfo)
}

func GetServiceAccount(b []byte) (*cloud_providers.ServiceAccount, error) {
	cloudInfo, err := UnMarshalCloudInfo(b)
	if err != nil {
		return nil, err
	}

	return &cloudInfo.ServiceAccount, nil
}

func GetDockerCredential(b []byte) (*cloud_providers.DockerRegistryCredentials, error) {
	cloudInfo, err := UnMarshalCloudInfo(b)
	if err != nil {
		return nil, err
	}
	return &cloudInfo.DockerRegistryCredentials, nil
}

func UnMarshalCloudInfo(b []byte) (*cloud_providers.RcInfo, error) {
	var info cloud_providers.RcInfo
	err := json.Unmarshal(b, &info)
	if err != nil {
		return nil, err
	}

	switch info.CloudProvider {
	case cloud_providers.Onprem:
		{
			out, err := json.Marshal(info.RcInfrastructure)
			if err != nil {
				return nil, err
			}
			var onpremProvider onprem.Provider
			err = json.Unmarshal(out, &onpremProvider)
			if err != nil {
				return nil, err
			}
			info.RcInfrastructure = onpremProvider
			return &info, nil
		}
	case cloud_providers.Azure:
		{
			out, err := json.Marshal(info.RcInfrastructure)
			if err != nil {
				return nil, err
			}
			var azureProvider azure.AzureProvider
			err = json.Unmarshal(out, &azureProvider)
			if err != nil {
				return nil, err
			}
			info.RcInfrastructure = azureProvider
			return &info, nil
		}
	case cloud_providers.AWS:
		{
			out, err := json.Marshal(info.RcInfrastructure)
			if err != nil {
				return nil, err
			}
			var awsProvider aws.Infrastructure
			err = json.Unmarshal(out, &awsProvider)
			if err != nil {
				return nil, err
			}
			info.RcInfrastructure = awsProvider
			return &info, nil
		}
	default:
		return nil, errors.New("unknown cloud provider")
	}
}

func GetControllerIp(b []byte) ([]string, error) {
	cloudInfo, err := UnMarshalCloudInfo(b)
	if err != nil {
		return nil, err
	}
	return cloudInfo.ControllerIPs, nil
}
