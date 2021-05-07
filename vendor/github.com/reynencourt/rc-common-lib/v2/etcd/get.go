package etcd

import (
	"context"

	"github.com/reynencourt/rc-common-lib/v2/state"
)

func (c *Client) Get(ctx context.Context, key string) (*state.Entity, error) {
	r, err := c.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if r.Count == 0 {
		return nil, ErrNoKeysFound
	}
	kv := r.Kvs[0]
	return &state.Entity{
		Key:  string(kv.Key),
		Data: kv.Value,
	}, nil
}
