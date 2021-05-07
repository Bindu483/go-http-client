package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2ListStorageClass = "V2ListStorageClass"

var V2ListStorageClassRoute = Route{
	Path:   "/api/v2/storageclass",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V2AddStorageClass = "V2AddStorageClass"

var V2AddStorageClassRoute = Route{
	Path:   "/api/v2/storageclass",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2GetStorageClassById = "V2GetStorageClassById"

var V2GetStorageClassByIdRoute = Route{
	Path:   "/api/v2/storageclass/{id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2UpdateStorageClass = "V2UpdateStorageClass"

var V2UpdateStorageClassRoute = Route{
	Path:   "/api/v2/storageclass/{id}",
	Method: http.MethodPut,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2DeleteStorageClass = "V2DeleteStorageClass"

var V2DeleteStorageClassRoute = Route{
	Path:   "/api/v2/storageclass/{id}",
	Method: http.MethodDelete,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V2StorageClassPVCs = "V2StorageClassPVCs"

var V2StorageClassPVCsRoute = Route{
	Path:   "/api/v2/storageclass/pvcs/{deployment_id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}
