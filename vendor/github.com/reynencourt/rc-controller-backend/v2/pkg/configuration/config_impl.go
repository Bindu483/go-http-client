package configuration

import (
	"context"
	"flag"
	"github.com/reynencourt/rc-common-lib/v2/shared_state"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ianschenck/envflag"
	"github.com/reynencourt/rc-common-lib/v2/cloud_info"
	rcEtcd "github.com/reynencourt/rc-common-lib/v2/etcd"
	"github.com/sirupsen/logrus"

	"github.com/reynencourt/rc-controller-backend/v2/pkg/controlleretcd"
	deploymentManager "github.com/reynencourt/rc-deployment-manager/v2/proto"
	resourceManager "github.com/reynencourt/rc-resource-manager/v2/proto"
	"google.golang.org/grpc"
)

type Manager struct {
	Router                   *mux.Router
	RcManager                resourceManager.ResourceManagerClient
	DeploymentManager        deploymentManager.DeploymentManagerClient
	EtcdClient               *controlleretcd.EtcdClientV2
	SharedState              *shared_state.Client
	NodeName                 string
	NodeIP                   string
	ServiceAccount           string
	DebugMode                bool
	Port                     string
	EnableHttps              bool
	UiHost                   string
	UMSHost                  string
	LawfirmHubURL            string
	BaseURL                  string
	GlusterMetricsConfigFile string
}

func New(r *mux.Router) *Manager {

	port := envflag.String("PORT", ":8888", "port for the service")
	rsManager := envflag.String("RESOURCE_MANAGER", "localhost:9090", "resource manager location")
	dplyManager := envflag.String("DEPLOYMENT_MANAGER", "localhost:9091", "resource manager location")
	etcdAddr := envflag.String("ETCD_ADDR", "http://localhost:2379", "location where the hosts file is backed up")
	nodeName := envflag.String("NODE_NAME", "", "node name of the machine")
	nodeIP := envflag.String("NODE_IP", "", "node ip of the machine")
	debug := envflag.Bool("DEBUG", false, "debug mode enabled")
	enableHTTPS := envflag.Bool("ENABLE_HTTPS", false, "enable redirect to https")
	uiHost := envflag.String("UI_HOST", "localhost:8000", "enable redirect to https")
	lawfirmService := envflag.String("LF_SERVICE", "https://lawfirms.rcplatform.io", "enable redirect to https")
	overrideBaseURL := envflag.String("BASE_URL", "", "enable redirect to https")
	glusterMetricsConfigFile := envflag.String("GLUSTER_METRICS_CONFIG", "", "gluster metrics toml config file")

	flag.Parse()
	envflag.Parse()

	if *nodeName == "" {
		logrus.Fatal("node name cannot be empty")
	}

	if *nodeIP == "" {
		logrus.Fatal("node name cannot be empty")
	}

	if *uiHost == "" {
		logrus.Fatal("ui host not specified")
	}

	var logger = logrus.
		WithField("port", *port).
		WithField("node_name", nodeName).
		WithField("node_ip", nodeIP).
		WithField("node_name", *nodeName).
		WithField("resource_manager", *rsManager).
		WithField("deployment_manager", *dplyManager)

	if strings.TrimSpace(*etcdAddr) == "" {
		logrus.Fatal("etcd address is not set")
	}

	urls := strings.Split(*etcdAddr, ",")
	etcdClient, err := rcEtcd.New(urls, rcEtcd.ClientPEMPath, rcEtcd.ClientKeyPath, rcEtcd.TrustedCAPath)
	if err != nil {
		logrus.WithError(err).Fatal("error while  connecting to etcd")
	}

	ctrlEtcd := controlleretcd.NewClientV2(urls, *nodeName, etcdClient)
	sharedState := shared_state.New(etcdClient)
	//etcdClient, err := rc_etcd.New(strings.Split(*etcdAddr, ","))
	//if err != nil {
	//	logrus.WithError(err).Fatal("error while  connecting to etcd")
	//}

	cloudinfoRaw, err := sharedState.GetCloudInfo(context.Background())
	if err != nil {
		logrus.Fatal("cloud config is not set yet")
	}

	cloudInfo, err := cloud_info.UnMarshalCloudInfo(cloudinfoRaw)
	if err != nil {
		logrus.Fatal("cloud config could not be parsed")
	}

	rcManagerConn, err := grpc.Dial(*rsManager, grpc.WithInsecure())
	if err != nil {
		logger.WithError(err).Fatal("error while connecting of Resource Manager")
	}

	deploymentManagerConn, err := grpc.Dial(*dplyManager, grpc.WithInsecure())
	if err != nil {
		logger.WithError(err).Fatal("error while connecting to deployment manager")
	}

	if err := sharedState.UpdateNodeInfo(context.Background(), *nodeName, *nodeIP); err != nil {
		logger.Error("error while updating node info")
	}

	var baseURL string
	if *overrideBaseURL == "" {
		baseURL = cloudInfo.BaseURL
	} else {
		baseURL = *overrideBaseURL
	}

	return &Manager{
		Router:                   r,
		RcManager:                resourceManager.NewResourceManagerClient(rcManagerConn),
		DeploymentManager:        deploymentManager.NewDeploymentManagerClient(deploymentManagerConn),
		EtcdClient:               ctrlEtcd,
		SharedState:              sharedState,
		NodeName:                 *nodeName,
		NodeIP:                   *nodeIP,
		DebugMode:                *debug,
		Port:                     *port,
		EnableHttps:              *enableHTTPS,
		UiHost:                   *uiHost,
		LawfirmHubURL:            *lawfirmService,
		BaseURL:                  baseURL,
		GlusterMetricsConfigFile: *glusterMetricsConfigFile,
	}
}
