package charter

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/container"
)

//this will simply update container registry domain in container.json
// and write the file back to helm chart
func (c *HelmChart) UpdateContainerJson(domain string) error {
	for _, f := range c.Files {
		if f.Name == "container.json" {
			var containers []container.ContainerImageReference
			if err := json.Unmarshal(f.Data, &containers); err != nil {
				return err
			}
			if len(containers) < 1 {
				return ErrChartHasNoImages
			}
			for _, c := range containers {
				c.ContainerImage.Registry = domain
			}
			updatedData, err := json.Marshal(&containers)
			if err != nil {
				return err
			}
			f.Data = updatedData
			return nil
		}
	}
	return errors.New("container.json not found")
}

func (c *HelmChart) UpdateContainerJsonWithRepoPrefix(domain, repoPrefix string) error {
	for _, f := range c.Files {
		if f.Name == "container.json" {
			var containers []container.ContainerImageReference
			if err := json.Unmarshal(f.Data, &containers); err != nil {
				return err
			}
			if len(containers) < 1 {
				return ErrChartHasNoImages
			}
			for _, c := range containers {
				c.ContainerImage.Registry = domain
				if repoPrefix != "" {
					c.ContainerImage.Repository = fmt.Sprintf("%s/%s", repoPrefix, c.ContainerImage.Repository)
				}
			}
			updatedData, err := json.Marshal(&containers)
			if err != nil {
				return err
			}
			f.Data = updatedData
			return nil
		}
	}
	return errors.New("container.json not found")
}
