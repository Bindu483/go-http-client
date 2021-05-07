package shared_state

import (
	"context"
	"encoding/json"
	"github.com/reynencourt/rc-common-lib/v2/proto/cluster"
	"github.com/reynencourt/rc-common-lib/v2/state"
)

func (c *Client) GetKubeConfig(ctx context.Context, clusterId string) ([]byte, error) {
	resp, err := c.Get(ctx, GetClusterConfigKey(clusterId))
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) GetClusterInfo(ctx context.Context, clusterId string) (*cluster.ClusterInfo, error) {
	d, err := c.State.Get(ctx, GetClusterInfoKey(clusterId))
	if err != nil {
		return nil, err
	}
	var out cluster.ClusterInfo

	err = json.Unmarshal(d.Data, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *Client) SaveClusterInfo(ctx context.Context, ci *cluster.ClusterInfo) error {
	key := GetClusterInfoKey(ci.Id)

	d, err := json.Marshal(ci)
	if err != nil {
		return err
	}
	return c.Put(ctx, &state.Entity{
		Key:  key,
		Data: d,
	})
}
