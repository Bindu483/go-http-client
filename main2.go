package main

import (
	"bytes"
	"encoding/json"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/nodeoperations"
	resourceManager "github.com/reynencourt/rc-resource-manager/v2/proto"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	postBody, _ := json.Marshal(map[string]string{
		"username": "admin@priyolahiri.in",
		"password": "xKOSJufaudkOSn=EcRH1E_yq",
	})
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, "https://v212-upgrade.rcplatform.io/api/v2/auth_server/login", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json")
	client := http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}
	responseByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	loginResponseObj := loginResponse{}
	err = json.Unmarshal(responseByte, &loginResponseObj)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(loginResponseObj.Meta.Token)
	//req, err = http.NewRequest(http.MethodGet, "https://v212-upgrade.rcplatform.io/api/v2/solutions", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Add("X-Access-Token", loginResponseObj.Meta.Token)
	//client = http.Client{
	//	Timeout: time.Duration(5) * time.Second,
	//}
	//resp, err = client.Do(req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if resp.StatusCode != http.StatusOK {
	//	log.Fatal(resp.StatusCode)
	//}
	//responseByte, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(responseByte))
	//var listSolRes deploymentManager.ListSolutionsResponse
	//
	//err = json.Unmarshal(responseByte, &listSolRes)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(listSolRes.Apps[0].SolutionId)
	//var createClusterReq = resourceManager.CreateK8SClusterRequest{
	//	ProviderType: 2,
	//	K8SSpec: &resourceManager.K8SSpec{
	//		ClusterName:       "bindu",
	//		MasterInstanceIps: []string{"10.0.1.111"},
	//		WorkerInstanceIps: []string{"10.0.1.99","10.0.1.16","10.0.1.252"},
	//		EtcdInstanceIps:   []string{"10.0.1.4"},
	//		Template:          "{{clusterName}}-{{k8s}}-{{serialNo}}",
	//		DeployFalco:       false,
	//		ClusterId:         "bindu1998",
	//	},
	//	ProjectId: "2",
	//	CreatedBy: "admin@priyolahiri.in",
	//}
	//
	//postBody, err = json.Marshal(&createClusterReq)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//req, err = http.NewRequest(http.MethodPost, "https://v212-upgrade.rcplatform.io/api/v2/clusters", bytes.NewBuffer(postBody))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	//req.Header.Add("X-Access-Token", loginResponseObj.Meta.Token)
	//req.Header.Add("x-Project", "1")
	//resp, err = client.Do(req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if resp.StatusCode != http.StatusOK {
	//	log.Println(resp.StatusCode)
	//	responseByte,_=ioutil.ReadAll(resp.Body)
	//	log.Fatal(string(responseByte))
	//}
	var createAddNode = nodeoperations.AddNodeRequest{
		ClusterId: "three",
		K8SSpec: &resourceManager.AddNodeK8SSpec{

			WorkerInstanceIps: []string{"10.0.1.14"},
		},
		CreatedBy: "admin@priyolahiri.in",
	}
	postBody, err = json.Marshal(&createAddNode)
	if err != nil {
		log.Fatal(err)
	}
	req, err = http.NewRequest(http.MethodPost, "https://v212-upgrade.rcplatform.io/api/v2/clusters/three/nodes", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Access-Token", loginResponseObj.Meta.Token)
	req.Header.Add("x-Project", "1")
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		responseByte, _ = ioutil.ReadAll(resp.Body)
		log.Fatal(string(responseByte))
	}

}

type loginResponse struct {
	Meta MetaLoginResponse `json:"meta"`
}
type MetaLoginResponse struct {
	Token string `json:"token"`
}
