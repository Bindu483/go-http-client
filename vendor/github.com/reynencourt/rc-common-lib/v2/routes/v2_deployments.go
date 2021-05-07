package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2ListDeployments = "V2ListDeployments"

var V2ListDeploymentsRoute = Route{
	Path:   "/api/v2/deployments",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Basic,
	},
	EnableTracing:    true,
	TracingEventName: "ListDeployments",
}

const V2GetDeploymentDetails = "V2GetDeploymentDetails"

var V2GetDeploymentDetailsRoute = Route{
	Path:   "/api/v2/deployments/{deployment_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "DeploymentDetails",
}

const V2GetDeploymentRcSpec = "V2GetDeploymentRcSpec"

var V2GetDeploymentRcSpecRoute = Route{
	Path:   "/api/v2/deployments/{deployment_id}/rcspec",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "GetDeploymentRcSpec",
}

const V2ListDeploymentEvents = "V2ListDeploymentEvents"

var V2ListDeploymentEventsRoute = Route{
	Path:   "/api/v2/deployments/{deployment_id}/events",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "ListDeploymentEvents",
}

const V2DeployApp = "V2DeployApp"

var V2DeployAppRoute = Route{
	Path:   "/api/v2/deployments",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "DeployApp",
}

const V2RemoveDeployment = "V2RemoveDeployment"

var V2RemoveDeploymentRoute = Route{
	Path:   "/api/v2/deployments/{deployment_id}",
	Method: http.MethodDelete,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2IsDeploymentUnique = "V2IsDeploymentUnique"

var V2IsDeploymentUniqueRoute = Route{
	Path:   "/api/v2/deployment_unique/{deployment_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "IsDeploymentUnique",
}

const V2GetCompatibleVersions = "V2GetCompatibleVersions"

var V2GetCompatibleVersionsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodGet,
	Path:   "/api/v2/apps/compatibleversion/{solution_id}/{version}",
}
