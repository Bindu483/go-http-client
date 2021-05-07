package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1ClusterMetrics = "V1ClusterMetrics"

var V1ClusterMetricsRoute = Route{
	Path:   "/api/v1/clusterMetrics",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1NodeMetrics = "V1NodeMetrics"

var V1NodeMetricsRoute = Route{
	Path:   "/api/v1/nodeMetrics",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1MetricsAppMetrics = "V1MetricsAppMetrics"

var V1MetricsAppMetricsRoute = Route{
	Path:   "/api/v1/metrics/appMetrics",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1ClusterMetricsMemoryUtilization = "V1ClusterMetricsMemoryUtilization"

var V1ClusterMetricsMemoryUtilizationRoute = Route{
	Path:   "/api/v1/clusterMetrics/memoryUtilization",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1GetClusterAlerts = "V1GetClusterAlerts"

var V1GetClusterAlertsRoute = Route{
	Path:   "/api/v1/getClusterAlerts",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}
