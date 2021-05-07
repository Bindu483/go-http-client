package shared_state

import (
	"time"

	"github.com/reynencourt/rc-common-lib/v2/proto/deploy_app"
)

type DeployedAppMetadata struct {
	AppName        string                      `json:"app_name"`
	DeploymentUUID string                      `json:"deployment_uuid"`
	DeployedOn     time.Time                   `json:"deployed_on"`
	Status         deploy_app.DeploymentStatus `json:"status"`
	Version        string                      `json:"version"`
	ClusterId      string                      `json:"cluster_id"`
}

type DeployedApp struct {
	Name     string              `json:"name"`
	Metadata DeployedAppMetadata `json:"metadata"`
}

type HostEntry struct {
	IP       string
	HostName string
}

type Member struct {
	MemberId   uint64
	IsLeader   bool
	IsHealthy  bool
	MemberName string
}

type Members []*Member

type IngressMessage struct {
	Ingress     []string
	ClusterName string
	ClusterId   string
	SystemApps  []string
}

type ClusterConfig struct {
	ClusterName   string
	ClusterId     string
	ClusterConfig string
}
