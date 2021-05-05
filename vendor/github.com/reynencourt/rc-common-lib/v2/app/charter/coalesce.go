package charter

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/reynencourt/helm/v3/pkg/chart"
	"path/filepath"
	"strings"
)

func (c *RcHelmChart) Coalesce() error {
	//manifests := *c.DeploymentManifest.Chart
	rawChart := chart.Chart(*c.rawChart)
	//handle root template first
	for templateIndex, template := range c.rawChart.Templates {
		if v, ok := c.DeploymentManifest.Chart[template.Name]; ok {
			k8sManifests := make([]string, len(v))
			for manifestIndex, manifest := range v {
				d, err := serializeK8sObject(manifest)
				if err != nil {
					return err
				}
				k8sManifests[manifestIndex] = string(d)
			}
			c.rawChart.Templates[templateIndex].Data = []byte(strings.Join(k8sManifests, "\n---\n"))
		}
	}

	for _, dc := range rawChart.Dependencies() {
		err := c.handleChartDependency(fmt.Sprintf("charts/%s/", dc.Metadata.Name), dc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RcHelmChart) handleChartDependency(templatePrefix string, chart *chart.Chart) error {
	for templateIndex, template := range chart.Templates {
		templatePath := filepath.Join(templatePrefix, template.Name)
		if v, ok := c.DeploymentManifest.Chart[templatePath]; ok {
			k8sManifests := make([]string, len(v))
			for manifestIndex, manifest := range v {
				d, err := serializeK8sObject(manifest)
				if err != nil {
					return errors.Wrap(err, "error while handling dependent chart "+templatePath)
				}
				k8sManifests[manifestIndex] = string(d)
			}
			chart.Templates[templateIndex].Data = []byte(strings.Join(k8sManifests, "\n---\n"))
		}
	}
	for _, dc := range chart.Dependencies() {
		err := c.handleChartDependency(filepath.Join(templatePrefix, fmt.Sprintf("charts/%s", dc.Metadata.Name)), dc)
		if err != nil {
			return err
		}
	}
	return nil
}
