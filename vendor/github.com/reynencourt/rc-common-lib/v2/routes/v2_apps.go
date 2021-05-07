package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2SyncSolutions = "V2SyncSolutions"

var V2SyncSolutionsRoute = Route{
	Path:   "/api/v2/apps/sync",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "SyncSolutions",
}

const V2DownloadRelease = "V2DownloadRelease"

var V2DownloadReleaseRoute = Route{
	Path:   "/api/v2/apps/download",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}

const V2GetAppMetadata = "V2GetAppMetadata"

var V2GetAppMetadataRoute = Route{
	Path:   "/api/v2/apps/{solution_id}/{version}/metadata",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}

const V2GetAppsPvcs = "V2GetAppsPvcs"

var V2GetAppsPvcsRoute = Route{
	Path:   "/api/v2/apps/{solutionId}/{version}/pvcs",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}
