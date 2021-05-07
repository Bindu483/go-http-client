package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2CreateCluster = "V2CreateCluster"

var V2CreateClusterRoute = Route{
	Path:   "/api/v2/clusters",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "CreateCluster",
}

const V2ListCluster = "V2ListCluster"

var V2ListClusterRoute = Route{
	Path:   "/api/v2/clusters",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V2GetClusterDetails = "V2GetClusterDetails"

var V2GetClusterDetailsRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V2ListNodes = "V2ListNodes"

var V2ListNodesRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}/nodes",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V2AddNode = "V2AddNode"

var V2AddNodeRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}/nodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "AddNode",
}

const V2RemoveNode = "V2RemoveNode"

var V2RemoveNodeRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}/nodes",
	Method: http.MethodDelete,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "RemoveNode",
}

const V2CanRemoveNode = "V2CanRemoveNode"

var V2CanRemoveNodeRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}/canremovenodes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "CanRemoveNode",
}

const V2DeleteCluster = "V2DeleteCluster"

var V2DeleteClusterRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}",
	Method: http.MethodDelete,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "DeleteCluster",
}

const V2MoveCluster = "V2MoveCluster"

var V2MoveClusterRoute = Route{
	Path:   "/api/v2/clusters/{cluster_id}/move",
	Method: http.MethodPut,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	EnableTracing:    true,
	TracingEventName: "MoveCluster",
}

const V2GetApprovedInstanceSizes = "V2GetApprovedInstanceSizes"

var V2GetApprovedInstanceSizesRoute = Route{
	Path:   "/api/v2/approved_instance_size/{provider_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}
