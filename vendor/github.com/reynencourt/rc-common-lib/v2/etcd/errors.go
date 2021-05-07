package etcd

import "github.com/pkg/errors"

var (
	ErrNoKeysFound      = errors.New("no keys found")
	ErrUnknownItemFound = errors.New("unknown element found")
)
