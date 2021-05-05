package resourceManager

import (
	"github.com/pkg/errors"
	"time"
)

var ErrCouldNotParseClusterLockTime = errors.New("failed to parse cluster lock time")

func (x *ClusterLock) IsValid() (bool, error) {
	validTil, err := time.Parse(time.RFC3339Nano, x.ValidTil)
	if err != nil {
		return true, errors.Wrap(err, ErrCouldNotParseClusterLockTime.Error())
	}

	if time.Now().Before(validTil) {
		return true, nil
	}
	return false, nil
}
