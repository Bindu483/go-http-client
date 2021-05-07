package etcd

import (
	"context"
	recipe "github.com/velann21/etcd/contrib/recipes"
)

func (c *Client) DeQueue(_ context.Context, prefix string) ([]byte, error) {
	var q *recipe.Queue
	var err error
	if len(c.queue) == 0 {
		c.queue = make(map[string]*recipe.Queue)
	}

	q, ok := c.queue[prefix]
	if !ok {
		q, err = c.NewQueue(prefix)
		if err != nil {
			return nil, err
		}
		c.queue[prefix] = q
	}

	d, err := q.Dequeue()
	return []byte(d), err
}
