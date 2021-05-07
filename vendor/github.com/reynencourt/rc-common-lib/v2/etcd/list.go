package etcd

import (
	"context"

	"github.com/reynencourt/rc-common-lib/v2/state"
	v3 "github.com/velann21/etcd/clientv3"
)

func (c *Client) List(ctx context.Context, keyPrefix string) (state.ListResponse, error) {
	r, err := c.client.Get(ctx, keyPrefix, v3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var out []*state.Entity
	for _, kv := range r.Kvs {
		out = append(out, &state.Entity{
			Key:  string(kv.Key),
			Data: kv.Value,
		})
	}
	return out, nil
}
