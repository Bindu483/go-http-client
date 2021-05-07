package state

import "context"

type Entity struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}

type ListResponse []*Entity

type State interface {
	Get(ctx context.Context, key string) (*Entity, error)
	Put(ctx context.Context, entity *Entity) error
	Delete(ctx context.Context, key string) (*Entity, error)
	List(ctx context.Context, keyPrefix string) (ListResponse, error)
	ListKeys(ctx context.Context, keyPrefix string) ([]string, error)
	Enqueue(ctx context.Context, prefix string, data []byte) error
	DeQueue(ctx context.Context, prefix string) ([]byte, error)
}
