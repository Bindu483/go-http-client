package routes

import (
	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
	"net/http"
)

const V2ListQosIndicatorForSolution = "V2ListQosIndicatorForSolution"

var V2ListQosIndicatorForSolutionRoute = Route{
	Path:   "/api/v2/qos/{solution_id}/{version}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2GetQosIndicatorForUpgrade = "V2GetQosIndicatorForUpgrade"

var V2GetQosIndicatorForUpgradeRoute = Route{
	Path:   "/api/v2/upgrade/qos/{deployment_id}/{version}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}
