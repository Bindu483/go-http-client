package shared_state

import "fmt"

const (
	KeyClusterLock              = "/v2/cluster/lock/%s"
	KeyPrivateKey               = "/privateKey"
	KeyCloudInfo                = "/cloudinfo"
	KeyControllerNodeInfoPrefix = "/nodeInfo"
	KeyControllerNodeInfo       = "/nodeInfo/%s"
	KeyHostEntries              = "/host-entries"
)

const (
	KeyClusterConfig             = "/v2/cluster/config/%s"
	KeyClusterInfo               = "/v2/cluster/info/%s"
	KeyClusterPrefix             = "/v2/cluster"
	KeyClusterResourceInfo       = "/v2/cluster/resource_info/%s"
	KeyClusterResourceInfoPrefix = `/v2/cluster/resource_info`
	KeyClusterConfigPrefix       = "/v2/cluster/config"
	KeyClusterModelPrefix        = "/v2/cluster/info"
)

func GetClusterLockKey(clusterId string) string {
	return fmt.Sprintf(KeyClusterLock, clusterId)
}

func GetControllerNodeInfoKey(nodeName string) string {
	return fmt.Sprintf(KeyControllerNodeInfo, nodeName)
}

func GetClusterConfigKey(clusterId string) string {
	return fmt.Sprintf(KeyClusterConfig, clusterId)
}

func GetClusterInfoKey(clusterId string) string {
	return fmt.Sprintf(KeyClusterInfo, clusterId)
}

func GetClusterResourceInfoKey(clusterId string) string {
	return fmt.Sprintf(KeyClusterResourceInfo, clusterId)
}
