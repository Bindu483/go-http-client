package charter

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/helm/v3/pkg/chartutil"
	"github.com/reynencourt/helm/v3/pkg/strvals"
	"strings"
)

func (c *HelmChart) CoalesceDeploymentValues(helmValues map[string]interface{}) (chartutil.Values, error) {
	rawChart := chart.Chart(*c)
	base := map[string]interface{}{}
	var valuesFromSpec []string
	for k, v := range helmValues {
		if v == "" {
			continue
		}
		valuesFromSpec = append(valuesFromSpec, fmt.Sprintf("%s=%s", k, v))
	}
	if err := strvals.ParseInto(strings.Join(valuesFromSpec, ","), base); err != nil {
		return nil, err
	}
	specValues, err := yaml.Marshal(&base)
	if err != nil {
		return nil, err
	}
	vals, err := chartutil.ReadValues(specValues)
	if err != nil {
		return nil, err
	}
	return chartutil.CoalesceValues(&rawChart, vals)
}
