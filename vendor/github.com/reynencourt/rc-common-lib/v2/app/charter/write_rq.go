package charter

import (
	"github.com/reynencourt/helm/v3/pkg/chart"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func (c *RcHelmChart) WriteResourceQuotaSpec(rq *corev1.ResourceQuota) error {
	data, err := yaml.Marshal(rq)
	if err != nil {
		return err
	}
	c.rawChart.Templates = append(c.rawChart.Templates, &chart.File{
		Name: "templates/rc_resource_quota.yaml",
		Data: data,
	})
	return nil
}
