package deploymentManager

import (
	"github.com/google/uuid"
	"github.com/reynencourt/rc-common-lib/v2/platform_variables"
)

func GetFakePlatformEnvVariables() map[string]string {
	randomID := uuid.New().String()
	return map[string]string{
		platform_variables.ClusterNameKey:            "default",
		platform_variables.DeploymentNameKey:         randomID,
		platform_variables.DeploymentUUIDKey:         randomID,
		platform_variables.PrimaryFQDNKey:            "app.default.fqdn.com",
		platform_variables.BaseDomainKey:             "default.fqdn.com",
		platform_variables.ContainerRegistryURLKey:   "https://someurl.com",
		platform_variables.PrimaryRWOStorageClassKey: "rc-storage",
		platform_variables.PrimaryRWXStorageClassKey: "rc-fs-storage",
		platform_variables.LawfirmNameKey:            "default",
	}
}

func AppendPlatformEnvVariables(a map[string]interface{}, b map[string]string) map[string]interface{} {
	t := a
	for k, v := range b {
		t[k] = v
	}
	return t
}
