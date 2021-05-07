package main

import (
	"github.com/Bindu483/go-http-client/auth"
	"github.com/Bindu483/go-http-client/cluster"
	rmCluster "github.com/reynencourt/rc-common-lib/v2/proto/cluster"
	"github.com/sirupsen/logrus"
)

func main() {

	loginSvc := &auth.Service{BaseUrl: "https://v212-upgrade.rcplatform.io"}

	authRes, err := loginSvc.Login()
	if err != nil {
		logrus.Fatalf("failed to login err:%s", err.Error())
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

	clusterSvc := &cluster.Service{
		BaseUrl:    "https://v212-upgrade.rcplatform.io",
		JWT:        authRes.JWT,
		ClusterId:  "testone",
		UserEmail:  authRes.Email,
		ProjectId:  "1",
		ProviderId: rmCluster.ProviderType_OnPrem,
	}
	err = clusterSvc.CreateCluster([]string{""}, []string{""}, []string{""})
	if err != nil {
		logrus.Fatalf("failed to create cluster err:%s", err.Error())
	}

	//var createAddNode = nodeoperations.AddNodeRequest{
	//	CreatedBy: "admin@priyolahiri.in",
	//	WorkerIps: []string{""},
	//}
	//postBody, err = json.Marshal(&createAddNode)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req, err = http.NewRequest(http.MethodPost, "https://v212-upgrade.rcplatform.io/api/v2/clusters/three/nodes", bytes.NewBuffer(postBody))
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
	//
	//if resp.StatusCode != http.StatusOK {
	//	log.Println(resp.StatusCode)
	//	responseByte, _ = ioutil.ReadAll(resp.Body)
	//	log.Fatal(string(responseByte))
	//}

}
