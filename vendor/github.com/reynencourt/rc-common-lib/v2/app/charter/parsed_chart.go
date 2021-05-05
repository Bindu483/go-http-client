package charter

import (
	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/helm/v3/pkg/chartutil"
	"github.com/reynencourt/rc-common-lib/v2/proto/rc_spec"
)

type ParsedChart struct {
	Chart  *chart.Chart
	RcSpec *rc_spec.RCSpec
}

func (c *ParsedChart) Save(outDir string) (string, error) {
	return chartutil.Save(c.Chart, outDir)
}
