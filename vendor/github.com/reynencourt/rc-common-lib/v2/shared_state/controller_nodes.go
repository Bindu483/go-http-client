package shared_state

import (
	"context"
	"fmt"
	"strings"

	"github.com/reynencourt/rc-common-lib/v2/etcd"
	"github.com/reynencourt/rc-common-lib/v2/state"
)

type Node struct {
	NodeName string
	NodeIP   string
}

func (c *Client) GetControllerNodes(ctx context.Context) ([]Node, error) {
	resp, err := c.List(ctx, KeyControllerNodeInfoPrefix)
	if err != nil {
		return nil, err
	}
	var nodes []Node
	for _, entity := range resp {
		keysTokens := strings.Split(entity.Key, "/")
		if len(keysTokens) != 3 {
			return nil, etcd.ErrUnknownItemFound
		}
		nodes = append(nodes, Node{
			NodeName: keysTokens[2],
			NodeIP:   string(entity.Data),
		})
	}
	return nodes, nil
}

func (c *Client) GetControllerNodeIp(ctx context.Context, nodeName string) (string, error) {
	k := fmt.Sprintf(KeyControllerNodeInfo, nodeName)
	resp, err := c.Get(ctx, k)
	if err != nil {
		return "", err
	}
	return string(resp.Data), nil
}

func (c *Client) UpdateNodeInfo(ctx context.Context, nodeName string, nodeIP string) error {
	e := &state.Entity{
		Key:  fmt.Sprintf(KeyControllerNodeInfo, nodeName),
		Data: []byte(nodeIP),
	}
	err := c.Put(ctx, e)
	if err != nil {
		return err
	}
	return nil
}
