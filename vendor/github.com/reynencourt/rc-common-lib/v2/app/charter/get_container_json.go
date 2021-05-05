package charter

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/container"
)

func (c *HelmChart) GetContainerJSON() ([]container.ContainerImageReference, error) {
	var containers []container.ContainerImageReference
	for _, f := range c.Files {
		if f.Name == "container.json" {
			if err := json.Unmarshal(f.Data, &containers); err != nil {
				return nil, errors.Wrap(err, "failed to unmarshall container.json")
			} else {
				return containers, nil
			}
		}
	}
	return nil, ErrContainerJsonNotFound
}
