package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1ListCluster = "V1ListCluster"

var V1ListClusterRoute = Route{
	Path:   "/api/v1/listCluster",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1ListNodes = "V1ListNodes"

var V1ListNodesRoute = Route{
	Path:   "/api/v1/listNodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1ListClusterWithoutNodes = "V1ListClusterWithoutNodes"

var V1ListClusterWithoutNodesRoute = Route{
	Path:   "/api/v1/listClusterWithoutNodes",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1CreateResource = "V1CreateResource"

var V1CreateResourceRoute = Route{
	Path:   "/api/v1/createResource",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1CreateResources = "V1CreateResources"

var V1CreateResourcesRoute = Route{
	Path:   "/api/v1/createResources",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "CreateResources",
}

const V1AddNodes = "V1AddNodes"

var V1AddNodesRoute = Route{
	Path:   "/api/v1/addNodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "AddNode",
}

const V1RemoveNodes = "V1RemoveNodes"

var V1RemoveNodesRoute = Route{
	Path:   "/api/v1/removeNodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "RemoveNode",
}

const V1MoveCluster = "V1MoveCluster"

var V1MoveClusterRoute = Route{
	Path:   "/api/v1/moveCluster",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1RemoveCluster = "V1RemoveCluster"

var V1RemoveClusterRoute = Route{
	Path:   "/api/v1/removeCluster",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "DeleteCluster",
}

const V1UpgradeCluster = "V1UpgradeCluster"

var V1UpgradeClusterRoute = Route{
	Path:   "/api/v1/upgradeCluster",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ClusterManagement,
	},
}

const V1CanRemoveNodes = "V1CanRemoveNodes"

var V1CanRemoveNodesRoute = Route{
	Path:   "/api/v1/canRemoveNodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

//
//const V1GetCephDashboardPassword = "V1GetCephDashboardPassword"
//
//var V1GetCephDashboardPasswordRoute = Route{
//	Path:   "/api/v1/getCephDashboardPassword",
//	Method: http.MethodPost,
//	Permissions: []permission.Permission{
//		permission.ProjectAdmin,
//	},
//}
