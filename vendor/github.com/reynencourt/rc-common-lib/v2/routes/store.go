package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1Content = "V1Content"

var V1ContentRoute = Route{
	Path:     "/api/v1/content",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ViewStoreSolutions,
	},
}

const V2Content = "V2Content"

var V2ContentRoute = Route{
	Path:   "/api/v2/content",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ViewStoreSolutions,
	},
}

const V1Solutions = "V1Solutions"

var V1SolutionsRoute = Route{
	Path:     "/api/v1/solutions",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ViewStoreSolutions,
	},
}

const V1StoreBroadcast = "V1StoreBroadcast"

var V1StoreBroadcastRoute = Route{
	Path:     "/api/v1/store/broadcast",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ViewStoreSolutions,
	},
}

const V1ContentSync = "V1ContentSync"

var V1ContentSyncRoute = Route{
	Path:   "/api/v1/content/sync",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.StoreSyncManager,
	},
}

const V1TestDriveSolutions = "V1TestDriveSolutions"

var V1TestDriveSolutionsRoute = Route{
	Path:     "/api/v1/testdrive/solutions",
	Method:   http.MethodGet,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1TestDriveProducts = "V1TestDriveProducts"

var V1TestDriveProductsRoute = Route{
	Path:     "/api/v1/testdrive/products",
	Method:   http.MethodGet,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1TestDriveStatus = "V1TestDriveStatus"

var V1TestDriveStatusRoute = Route{
	Path:     "/api/v1/testdrive/status",
	Method:   http.MethodGet,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1TestDriveExists = "V1TestDriveExists"

var V1TestDriveExistsRoute = Route{
	Path:     "/api/v1/testdrive/exists",
	Method:   http.MethodPost,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ViewStoreSolutions,
	},
}

const V1TestDriveInitiate = "V1TestDriveInitiate"

var V1TestDriveInitiateRoute = Route{
	Path:     "/api/v1/testdrive/initiate",
	Method:   http.MethodPost,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.TestDriveAdmin,
	},
}

const V1TestDriveLaunchPad = "V1TestDriveLaunchPad"

var V1TestDriveLaunchPadRoute = Route{
	Path:   "/api/v1/testdrive/all",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}
