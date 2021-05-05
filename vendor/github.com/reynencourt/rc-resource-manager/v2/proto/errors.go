package resourceManager

import "github.com/pkg/errors"

var ErrClusterWithSameIdExist = errors.New("cluster with same id already exist")
