package platform_variables

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/reynencourt/rc-common-lib/v2/proto/rc_spec"
)

const (
	ClusterNameKey            = "rc.deployment.cluster_name"
	DeploymentNameKey         = "rc.deployment.name"
	DeploymentUUIDKey         = "rc.deployment.uuid"
	PrimaryFQDNKey            = "rc.deployment.primary_fqdn"
	BaseDomainKey             = "rc.deployment.base_domain"
	ContainerRegistryURLKey   = "rc.deployment.container_registry_url"
	PrimaryRWOStorageClassKey = "rc.storage.primary_rwo_storage_class"
	PrimaryRWXStorageClassKey = "rc.storage.primary_rwx_storage_class"
	LawfirmNameKey            = "rc.lawfirm_name"
)

func GetFakePlatformEnvVariables() map[string]string {
	randomID := uuid.New().String()
	return map[string]string{
		ClusterNameKey:            "default",
		DeploymentNameKey:         randomID,
		DeploymentUUIDKey:         randomID,
		PrimaryFQDNKey:            "app.default.fqdn.com",
		BaseDomainKey:             "default.fqdn.com",
		ContainerRegistryURLKey:   "https://someurl.com",
		PrimaryRWOStorageClassKey: "rc-storage",
		PrimaryRWXStorageClassKey: "rc-fs-storage",
		LawfirmNameKey:            "default",
	}
}

func SetPlatformEnvVariables(deploymentName, clusterName, deploymentUUID, baseUrl, lFName, containerRegistryDomain string, services *rc_spec.RCServices) map[string]string {
	out := make(map[string]string)
	primaryFQDN, fqdn := getFQDN(deploymentName, clusterName, baseUrl, services)
	out[ClusterNameKey] = clusterName
	out[DeploymentNameKey] = deploymentName
	out[DeploymentUUIDKey] = deploymentUUID
	out[PrimaryFQDNKey] = primaryFQDN
	out[BaseDomainKey] = fqdn
	out[ContainerRegistryURLKey] = containerRegistryDomain
	out[PrimaryRWOStorageClassKey] = "rc-storage"
	out[PrimaryRWXStorageClassKey] = "rc-fs-storage"
	out[LawfirmNameKey] = lFName
	return out
}

func getFQDN(deploymentName, clusterName, rootDomain string, s *rc_spec.RCServices) (string, string) {
	fqdn := fmt.Sprintf("%s-%s.%s", clusterName, deploymentName, rootDomain)
	for _, service := range s.Expose {
		if service.Primary {
			if strings.TrimSpace(service.DomainPrefix) != "" {
				return fmt.Sprintf("%s.%s", service.DomainPrefix, fqdn), fqdn
			}
		}
	}
	return fqdn, fqdn
}
