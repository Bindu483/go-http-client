package cluster

import (
	"sort"
	"time"

	"github.com/reynencourt/rc-common-lib/v2/ops/sshremoteexec"
)

type Nodes []*Node
type Pred func(node *Node) bool

func (nodes Nodes) Filter(pred Pred) Nodes {
	var out Nodes

	for _, node := range nodes {
		if pred(node) {
			out = append(out, node)
		}
	}
	return out
}

func (nodes Nodes) GetNodeNames() []string {
	var out []string

	for _, n := range nodes {
		out = append(out, n.HostName)
	}
	return out
}

func (n *Node) IsReachable(keyLocation string) error {
	client := sshremoteexec.SshConnection{
		User:        n.User,
		Ip:          n.Ip,
		KeyLocation: keyLocation,
		Timeout:     5 * time.Second,
	}
	return client.IsMachineReachable()
}

func (nodes Nodes) Len() int           { return len(nodes) }
func (nodes Nodes) Less(i, j int) bool { return nodes[i].Number > nodes[j].Number }
func (nodes Nodes) Swap(i, j int)      { nodes[i], nodes[j] = nodes[j], nodes[i] }

func GenerateNextIndexForNodes(tag NodeType, nodes Nodes) int32 {
	var nodeVariety Nodes
	for _, node := range nodes {
		if node.Type == tag {
			nodeVariety = append(nodeVariety, node)
		}
	}
	if len(nodeVariety) == 0 {
		return 0
	}
	sort.Sort(nodeVariety)
	return nodeVariety[0].Number + 1
}

func GetTag(tag string) string {
	if tag == "master" {
		return "k8sm"
	} else if tag == "worker" {
		return "k8sw"
	} else {
		return tag
	}
}
