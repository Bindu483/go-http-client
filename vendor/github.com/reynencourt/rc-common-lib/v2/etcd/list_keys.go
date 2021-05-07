package etcd

import (
	"context"
	v3 "github.com/velann21/etcd/clientv3"
)

func (c Client) ListKeys(ctx context.Context, keyPrefix string) ([]string, error) {
	r, err := c.client.Get(ctx, keyPrefix, v3.WithPrefix(), v3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	if r.Count == 0 {
		return nil, ErrNoKeysFound
	}
	var out []string

	for _, kv := range r.Kvs {
		out = append(out, string(kv.Key))
	}
	return out, nil
}
