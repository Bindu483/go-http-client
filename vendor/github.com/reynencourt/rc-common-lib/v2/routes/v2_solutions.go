package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2ListSolutions = "V2ListSolutions"

var V2ListSolutionsRoute = Route{
	Path:   "/api/v2/solutions",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "ListSolutions",
}

const V2GetSolution = "V2GetSolution"

var V2GetSolutionRoute = Route{
	Path:   "/api/v2/solutions/{solution_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "GetSolution",
}

const V2ListSolutionVersions = "V2ListSolutionVersions"

var V2ListSolutionVersionsRoute = Route{
	Path:   "/api/v2/solutions/{solution_id}/versions",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "ListSolutionVersions",
}

const V2GetAppReleaseRcSpec = "V2GetAppReleaseRcSpec"

var V2GetAppReleaseRcSpecRoute = Route{
	Path:   "/api/v2/solutions/{solution_id}/{solution_version}/rcspec",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "GetAppReleaseRcSpec",
}
