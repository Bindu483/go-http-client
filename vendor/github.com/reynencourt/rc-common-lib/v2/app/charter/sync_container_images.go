package charter

import (
	"encoding/json"
	"fmt"

	"github.com/reynencourt/rc-common-lib/v2/proto/container"

	"github.com/pkg/errors"
)

var (
	ErrContainerJsonNotFound = errors.New("containers.json not found")
	ErrChartHasNoImages      = errors.New("chart has no container images")
)

//TODO: split responsibility. this should only update the chart with new domain
func (c *RcHelmChart) SyncContainerImages(domain string) error {
	for _, f := range c.rawChart.Files {
		if f.Name == "container.json" {
			var containers []container.ContainerImageReference
			if err := json.Unmarshal(f.Data, &containers); err != nil {
				return err
			}
			if len(containers) < 1 {
				return ErrChartHasNoImages
			}
			for _, c := range containers {
				err := c.ContainerImage.Pull()
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("failed to pull container %s", c.ContainerImage.ToString()))
				}
				d := *c.ContainerImage
				d.Registry = domain
				err = c.ContainerImage.TagImage(&d)
				if err != nil {
					return errors.Wrap(err, "failed to tag images")
				}
				err = d.Push()
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("failed to push container %s", c.ContainerImage.ToString()))
				}
			}
			return nil
		}
	}
	return ErrContainerJsonNotFound
}
