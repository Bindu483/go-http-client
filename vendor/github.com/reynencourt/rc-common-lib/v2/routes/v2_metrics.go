package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2Metrics = "V2Metrics"

var V2MetricsRoute = Route{
	Path:   "/api/v2/metrics",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}
