package charter

import (
	"encoding/json"

	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/rc-common-lib/v2/proto/container"
)

func (c *RcHelmChart) WriteContainerJSON(containers []container.ContainerImageReference) error {
	data, err := json.Marshal(containers)
	if err != nil {
		return err
	}
	c.rawChart.Files = append(c.rawChart.Files, &chart.File{
		Name: "container.json",
		Data: data,
	})
	return nil
}
