package etcd

import (
	"context"

	"github.com/reynencourt/rc-common-lib/v2/state"
	v3 "github.com/velann21/etcd/clientv3"
)

func (c *Client) Delete(ctx context.Context, key string) (*state.Entity, error) {
	r, err := c.client.Delete(ctx, key, v3.WithPrevKV())
	if err != nil {
		return nil, err
	}
	if r.Deleted == 0 {
		return nil, ErrNoKeysFound
	}
	kv := r.PrevKvs[0]
	return &state.Entity{
		Key:  string(kv.Key),
		Data: kv.Value,
	}, nil
}
