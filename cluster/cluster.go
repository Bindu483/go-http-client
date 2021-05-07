package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/Bindu483/go-http-client/utils"
	"github.com/reynencourt/rc-common-lib/v2/proto/cluster"
	resourceManager "github.com/reynencourt/rc-resource-manager/v2/proto"
	"log"
	"net/http"
)

type Service struct {
	BaseUrl    string               `json:"base_url"`
	JWT        string               `json:"jwt"`
	ClusterId  string               `json:"cluster_id"`
	UserEmail  string               `json:"user_email"`
	ProjectId  string               `json:"project_id"`
	ProviderId cluster.ProviderType `json:"provider_id"`
}

func (s *Service) CreateCluster(masterIps, etcdIps, workerIps []string) error {
	var createClusterReq = resourceManager.CreateK8SClusterRequest{
		ProviderType: s.ProviderId,
		K8SSpec: &resourceManager.K8SSpec{
			ClusterName:       s.ClusterId,
			MasterInstanceIps: masterIps,
			WorkerInstanceIps: workerIps,
			EtcdInstanceIps:   etcdIps,
			Template:          "{{clusterName}}-{{k8s}}-{{serialNo}}",
			DeployFalco:       false,
			ClusterId:         s.ClusterId,
		},
		ProjectId: s.ProjectId,
		CreatedBy: s.UserEmail,
	}

	postBody, err := json.Marshal(&createClusterReq)
	if err != nil {
		log.Fatal(err)
	}

	createClusterUrl := fmt.Sprintf("%s/api/v2/clusters", s.BaseUrl)

	headers := map[string]string{
		"Content-Type":   "application/json;charset=UTF-8",
		"Accept":         "application/json",
		"x-Project":      s.ProjectId,
		"X-Access-Token": s.JWT,
	}

	_, err = utils.MakeHTTPRequest(createClusterUrl, http.MethodPost, postBody, headers)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddNode(workerIps []string) error {
	return nil
}
