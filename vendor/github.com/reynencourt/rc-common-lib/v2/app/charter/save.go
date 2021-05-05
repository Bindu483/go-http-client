package charter

import (
	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/helm/v3/pkg/chartutil"
)

func (c *RcHelmChart) Save(outDir string) (string, error) {
	rawChart := chart.Chart(*c.rawChart)
	return chartutil.Save(&rawChart, outDir)
}

func (c *HelmChart) SetChartName(chartName string) {
	c.Metadata.Name = chartName
}

func (c *HelmChart) Save(outDir string) (string, error) {
	rawChart := chart.Chart(*c)
	return chartutil.Save(&rawChart, outDir)
}
