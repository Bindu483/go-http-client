package charter

import (
	"bytes"

	"github.com/reynencourt/helm/v3/pkg/chart/loader"

	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/rc-common-lib/v2/proto/rc_spec"
	"k8s.io/apimachinery/pkg/runtime"
)

type HelmChart chart.Chart
type HelmValues map[string]interface{}

type RcHelmChart struct {
	rawChart           *HelmChart            `json:"raw_chart"`
	DeploymentManifest *RcDeploymentManifest `json:"deployment_manifest"`
}

type RcDeploymentManifest struct {
	RcSpec *rc_spec.RCSpec `json:"rc_spec"`
	Chart  ParsedHelmChart `json:"k8s_manifests"`
}

type K8sObject runtime.Object

type RenderedChart map[string][]string

type K8sObjects []K8sObject

type ParsedHelmChart map[string]K8sObjects

func LoadChart(chartPath string) (*HelmChart, error) {
	c, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}
	rcChart := HelmChart(*c)
	return &rcChart, err
}

func LoadChartFromBytes(archive []byte) (*HelmChart, error) {
	c, err := loader.LoadArchive(bytes.NewReader(archive))
	if err != nil {
		return nil, err
	}
	rcChart := HelmChart(*c)
	return &rcChart, err
}

func (c *HelmChart) NewRcHelmChart(spec *rc_spec.RCSpec, parsedChart ParsedHelmChart) *RcHelmChart {
	return &RcHelmChart{
		rawChart: c,
		DeploymentManifest: &RcDeploymentManifest{
			RcSpec: spec,
			Chart:  parsedChart,
		},
	}
}
