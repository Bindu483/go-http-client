package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1ProjectViewerListApps = "V1ProjectViewerListApps"

var V1ProjectViewerListAppsRoute = Route{
	Path:   "/api/v1/projectviewer/listApps",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1ListApps = "V1ListApps"

var V1ListAppsRoute = Route{
	Path:   "/api/v1/listApps",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1GetApp = "V1GetApp"

var V1GetAppRoute = Route{
	Path:   "/api/v1/getApp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1DeployApp = "V1DeployApp"

var V1DeployAppRoute = Route{
	Path:   "/api/v1/deployApp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "DeployApps",
}

const V1SaveApp = "V1SaveApp"

var V1SaveAppRoute = Route{
	Path:   "/api/v1/saveApp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1ApplicationManagementListApps = "V1ApplicationManagementListApps"

var V1ApplicationManagementListAppsRoute = Route{
	Path:   "/api/v1/applicationmanagement/listApps",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}
