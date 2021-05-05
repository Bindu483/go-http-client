package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kube struct {
	clientset *kubernetes.Clientset
	config    []byte
}

const (
	NamespaceRcSystem   = "rc-system"
	ClusterNameLabelKey = "cluster_name"
	RcUser              = "reynen_court"
)

var (
	defaultLabels = map[string]string{
		"created_by": RcUser,
	}
	TimeoutSeconds int64 = 15
)

func New(config []byte) (*Kube, error) {
	c, err := clientcmd.NewClientConfigFromBytes(config)
	if err != nil {
		return nil, err
	}

	restConfig, err := c.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return &Kube{clientset: clientset, config: config}, nil
}

func NewFromKubeConfigFile(configPath string) (*Kube, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Kube{clientset: clientset}, nil
}

func GetDefaultLables(clusterName string) map[string]string {
	l := defaultLabels
	l[ClusterNameLabelKey] = clusterName
	return l
}

func (k *Kube) GetConfig() []byte {
	return k.config
}

func (k *Kube) GetK8sClientSet() *kubernetes.Clientset {
	return k.clientset
}
