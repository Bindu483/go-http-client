package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1LogsGetFalcoLogs = "V1LogsGetFalcoLogs"

var V1LogsGetFalcoLogsRoute = Route{
	Path:     "/api/v1/logs/getFalcoLogs",
	Method:   http.MethodPost,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}
