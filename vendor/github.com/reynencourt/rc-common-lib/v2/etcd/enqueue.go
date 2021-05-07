package etcd

import (
	"context"
	recipe "github.com/velann21/etcd/contrib/recipes"
)

func (c *Client) Enqueue(_ context.Context, prefix string, data []byte) error {
	var q *recipe.Queue
	var err error

	q, ok := c.queue[prefix]
	if !ok {
		q, err = c.NewQueue(prefix)
		if err != nil {
			return err
		}
		c.queue[prefix] = q
	}
	return q.Enqueue(string(data))
}
