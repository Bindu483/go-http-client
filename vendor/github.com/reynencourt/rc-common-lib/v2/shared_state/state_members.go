package shared_state

import (
	"context"
	"fmt"

	"github.com/reynencourt/rc-common-lib/v2/etcd"
	"github.com/sirupsen/logrus"
)

func (c *Client) GetMembers(ctx context.Context) (Members, error) {
	var outMap = make(map[string]*Member)

	typedEtcdClient, ok := c.GetState().(*etcd.Client)
	if !ok {
		panic("state cant be type casted to etcd")
	}

	members, err := typedEtcdClient.GetClient().MemberList(ctx)
	if err != nil {
		return nil, err
	}

	for _, m := range members.Members {
		outMap[m.Name] = &Member{
			MemberId:   m.ID,
			IsLeader:   false,
			IsHealthy:  false,
			MemberName: m.Name,
		}
	}

	nodes, err := c.GetControllerNodes(ctx)
	if err != nil {
		return nil, err
	}

	for _, n := range nodes {
		s, err := typedEtcdClient.GetClient().Status(ctx, fmt.Sprintf(`https://%s:2379`, n.NodeIP))
		if err != nil {
			logrus.WithError(err).Error("failed to get etcd node status fro %s", n.NodeIP)
			continue
		}
		outMap[n.NodeName].IsHealthy = true
		if outMap[n.NodeName].MemberId == s.Leader {
			outMap[n.NodeName].IsLeader = true
		}
	}
	var out []*Member

	for _, v := range outMap {
		out = append(out, v)
	}

	return out, nil
}

func (m Members) IsMaster(nodeName string) bool {
	for _, member := range m {
		if member.MemberName == nodeName {
			return member.IsLeader
		}
	}
	return false
}

func (m Members) FilterHealthyMembers() Members {
	var out []*Member
	for _, member := range m {
		if member.IsHealthy {
			out = append(out, member)
		}
	}
	return out
}
