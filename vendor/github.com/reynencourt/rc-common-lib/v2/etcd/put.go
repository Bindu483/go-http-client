package etcd

import (
	"context"

	"github.com/reynencourt/rc-common-lib/v2/state"
)

func (c *Client) Put(ctx context.Context, entity *state.Entity) error {
	_, err := c.client.Put(ctx, entity.Key, string(entity.Data))
	if err != nil {
		return err
	}
	return nil
}
