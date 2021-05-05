package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type NodeRole string

const (
	NodeRoleMaster     NodeRole = "master"
	NodeRoleWorker     NodeRole = "worker"
	NodeResourceCPU             = "cpu"
	NodeResourceMemory          = "memory"
)

type K8sNode struct {
	Name               string            `json:"name"`
	Taints             []v1.Taint        `json:"taints"`
	Labels             map[string]string `json:"labels"`
	InternalIP         string            `json:"internal_ip"`
	Capacity           v1.ResourceList   `json:"resources_allocated"`
	Allocatable        v1.ResourceList   `json:"resources_available"`
	Disk               v1.ResourceList   `json:"disk"`
	Ready              bool              `json:"ready"`
	HasMemoryPressure  bool              `json:"has_memory_pressure"`
	HasDiskPressure    bool              `json:"has_disk_pressure"`
	HasPIDPressure     bool              `json:"has_pid_pressure"`
	NetworkUnAvailable bool              `json:"network_un_available"`
	Role               NodeRole          `json:"role"`
}

type K8sNodes []K8sNode

type Predicate func(node K8sNode) bool

func (nodes K8sNodes) Filter(pred Predicate) K8sNodes {
	var filteredNodes = K8sNodes{}
	for _, node := range nodes {
		if pred(node) {
			filteredNodes = append(filteredNodes, node)
		}
	}
	return filteredNodes
}

func (k *Kube) GetNodes(ctx context.Context) (K8sNodes, error) {

	nList, err := k.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{TimeoutSeconds: &TimeoutSeconds})
	if err != nil {
		return nil, err
	}

	return typeCastNodeListToK8sNodes(nList), nil
}

func typeCastNodeListToK8sNodes(nList *v1.NodeList) K8sNodes {
	var nodes = K8sNodes{}
	for _, n := range nList.Items {
		node := K8sNode{
			Name:               n.Name,
			Taints:             n.Spec.Taints,
			Labels:             n.GetLabels(),
			InternalIP:         "",
			Capacity:           n.Status.Capacity,
			Allocatable:        n.Status.Allocatable,
			Disk:               n.Status.Allocatable,
			Ready:              false,
			HasMemoryPressure:  false,
			HasDiskPressure:    false,
			HasPIDPressure:     false,
			NetworkUnAvailable: false,
		}

		for k := range node.Labels {
			role := strings.Split(k, "/")
			if role[0] == "node-role.kubernetes.io" && role[1] == "master" {
				node.Role = NodeRoleMaster
			} else if role[0] == "node-role.kubernetes.io" && role[1] == "node" {
				node.Role = NodeRoleWorker
			}
		}
		for _, a := range n.Status.Addresses {
			if a.Type == v1.NodeInternalIP {
				node.InternalIP = a.Address
			}
		}

		for _, c := range n.Status.Conditions {
			if c.Type == v1.NodeReady && c.Status == v1.ConditionTrue {
				node.Ready = true
			}
			if c.Type == v1.NodeMemoryPressure && c.Status == v1.ConditionTrue {
				node.HasMemoryPressure = true
			}
			if c.Type == v1.NodeDiskPressure && c.Status == v1.ConditionTrue {
				node.HasDiskPressure = true
			}
			if c.Type == v1.NodePIDPressure && c.Status == v1.ConditionTrue {
				node.HasPIDPressure = true
			}
			if c.Type == v1.NodeNetworkUnavailable && c.Status == v1.ConditionTrue {
				node.NetworkUnAvailable = true
			}
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func (nodes K8sNodes) GetHealthyWorkers() K8sNodes {
	return nodes.Filter(func(node K8sNode) bool {
		return node.Ready
	})
}

func (nodes K8sNodes) GetSchedulableNodes() K8sNodes {
	return nodes.Filter(func(node K8sNode) bool {
		var nonSchedulable = true
		for _, t := range node.Taints {
			if t.Effect == v1.TaintEffectNoSchedule {
				nonSchedulable = false
				break
			}
		}
		return nonSchedulable
	})
}

func (nodes K8sNodes) GetNode(nodeName []string) K8sNodes {
	return nodes.Filter(func(node K8sNode) bool {
		for _, v := range nodeName {
			if node.Name == v {
				return true
			}
		}
		return false
	})
}
