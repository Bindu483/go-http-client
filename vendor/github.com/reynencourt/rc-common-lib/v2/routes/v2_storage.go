package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2ListStorage = "V2ListStorage"

var V2ListStorageRoute = Route{
	Path:   "/api/v2/storage",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2AddStorage = "V2AddStorage"

var V2AddStorageRoute = Route{
	Path:   "/api/v2/storage",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2GetStorageById = "V2GetStorageById"

var V2GetStorageByStorageIdRoute = Route{
	Path:   "/api/v2/storage/{id}",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2UpdateStorage = "V2UpdateStorage"

var V2UpdateStorageRoute = Route{
	Path:   "/api/v2/storage/{id}",
	Method: http.MethodPut,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2DeleteStorage = "V2DeleteStorage"

var V2DeleteStorageRoute = Route{
	Path:   "/api/v2/storage/{id}",
	Method: http.MethodDelete,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2UpdateStorageProjects = "V2UpdateStorageProjects"

var V2UpdateStorageProjectsRoute = Route{
	Path:   "/api/v2/storage/{id}/projects",
	Method: http.MethodPut,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}

const V2TestStorage = "V2TestStorage"

var V2TestStorageRoute = Route{
	Path:   "/api/v2/storage/test",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.StorageManagement,
	},
}
