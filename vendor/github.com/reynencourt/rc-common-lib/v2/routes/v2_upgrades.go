package routes

import (
	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
	"net/http"
)

const V2UpgradeApp = "V2UpgradeApp"

var V2UpgradeAppRoute = Route{
	Path:   "/api/v2/upgrade/{deployment_id}",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "UpgradeApp",
}

const V2UpdateConfig = "V2UpdateConfig"

var V2UpdateConfigRoute = Route{
	Path:   "/api/v2/update_config/{deployment_id}",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "UpdateConfig",
}
