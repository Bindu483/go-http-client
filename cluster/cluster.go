package cluster

type Service struct {
	BaseUrl    string `json:"base_url"`
	JWT        string `json:"jwt"`
	ClusterId  string `json:"cluster_id"`
	UserEmail  string `json:"user_email"`
	ProjectId  string `json:"project_id"`
	ProviderId string `json:"provider_id"`
}

func (s Service) CreateCluster(masterIps, etcdIps, workerIps []string) error {

}
