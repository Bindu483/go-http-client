package cloud_providers

import (
	"encoding/base64"
	"fmt"
)

type FirewallRule struct {
	Protocol string `json:"protocol"`
	FromPort int64  `json:"from_port"`
	ToPort   int64  `json:"to_port"`
	FromCIDR string `json:"from_cidr"`
	Type     int8   `json:"type"`
}

type FireWallRulesInput struct {
	Name  string          `json:"name"`
	Rules []*FirewallRule `json:"rules"`
}

type FireWallRuleOutput struct {
	Name           string `json:"name"`
	FirewallRuleId string `json:"firewall_rule_id"`
}

const IngressRule = 0
const EgressRule = 1

const (
	CloudProviderUnknown CloudProvider = 0
	AWS                  CloudProvider = 1
	Onprem               CloudProvider = 2
	Azure                CloudProvider = 3
	GCP                  CloudProvider = 4
)

type CloudProvider int

func (c CloudProvider) String() string {
	switch c {
	case AWS:
		return "AWS"
	case Onprem:
		return "Onprem"
	case Azure:
		return "Azure"
	case GCP:
		return "GCP"
	default:
		return "CloudProviderUnknown"
	}
}

func GetProvider(id int) CloudProvider {
	switch id {
	case 1:
		return AWS
	case 2:
		return Onprem
	case 3:
		return Azure
	case 4:
		return GCP
	default:
		return CloudProviderUnknown
	}
}

type ServiceAccount struct {
	Success     bool   `json:"success"`
	Token       string `json:"token"`
	Username    string `json:"username"`
	UserId      string `json:"userId"`
	LawfirmName string `json:"lawfirmName"`
}

type SMTPCredential struct {
	Success  bool   `json:"success"`
	Username string `json:"smtp_username"`
	Password string `json:"smtp_password"`
	Port     int    `json:"smtp_port"`
	Ssl      bool   `json:"smtp_ssl_tls"`
	Host     string `json:"smtp_host"`
	Sender   string `json:"smtp_sender"`
}

type DockerRegistryCredentials struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RegistryURL string `json:"registry"`
}

func (d *DockerRegistryCredentials) userPasswordCombination() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", d.Username, d.Password)))
}

func (d *DockerRegistryCredentials) GenerateDockerConfig() string {

	var config = `{
	"auths": {
   "%v" : {
			"auth": "%v"
    }},
	"HttpHeaders": {
		"User-Agent": "Docker-Client/18.06.3-ce (linux)"
	}
}`

	return fmt.Sprintf(config, d.RegistryURL, d.userPasswordCombination())
}

type RcInfo struct {
	CloudProvider             CloudProvider             `json:"cloud_provider"`
	AccessTokenUsed           string                    `json:"access_token"`
	RCLawfirmURL              string                    `json:"lawfirm_hub_url"`
	CreatorEmail              string                    `json:"creator_email"`
	ConsumerKey               string                    `json:"consumer_key"`
	DeploymentName            string                    `json:"deployment_name"`
	ControllerIPs             []string                  `json:"controller_ips"`
	HighAvailability          bool                      `json:"ha_enabled"`
	ServiceAccount            ServiceAccount            `json:"service_account"`
	RcDeploymentID            string                    `json:"deployment_id"`
	GeneratedPassword         string                    `json:"first_generated_password"`
	Version                   string                    `json:"version"`
	RcInfrastructure          interface{}               `json:"rc_infrastructure"`
	LawfirmName               string                    `json:"lawfirm"`
	SMTPCredential            SMTPCredential            `json:"smtp_credentials"`
	FirstName                 string                    `json:"first_name"`
	LastName                  string                    `json:"last_name"`
	RootUser                  string                    `json:"root_user"`
	RootPassword              string                    `json:"root_password"`
	DockerRegistryCredentials DockerRegistryCredentials `json:"docker_registry_credentials"`
	BaseURL                   string                    `json:"base_url"`
	OsType                    OsType                    `json:"os_type"`
}
