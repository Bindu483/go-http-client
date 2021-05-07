package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1GetStatus = "V1GetStatus"

var V1GetStatusRoute = Route{
	Path:   "/api/v1/getStatus",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1PlatformInfo = "V1PlatformInfo"

var V1PlatformInfoRoute = Route{
	Path:   "/api/v1/platformInfo",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1CloudConfigObject = "V1CloudConfigObject"

var V1CloudConfigObjectRoute = Route{
	Path:   "/api/v1/cloudConfigObject",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1GetApprovedInstanceSizes = "V1GetApprovedInstanceSizes"

var V1GetApprovedInstanceSizesRoute = Route{
	Path:   "/api/v1/getApprovedInstanceSizes",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1PlatformOperationsSmtp = "V1PlatformOperationsSmtp"

var V1PlatformOperationsSmtpRoute = Route{
	Path:   "/api/v1/platformoperations/smtp",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.PlatformConfiguration,
	},
}

const V1PlatformOperationsSmtpUpdate = "V1PlatformOperationsSmtpUpdate"

var V1PlatformOperationsSmtpUpdateRoute = Route{
	Path:   "/api/v1/platformoperations/smtp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.PlatformConfiguration,
	},
}

const V1PlatformOperationsSmtpTest = "V1PlatformOperationsSmtpTest"

var V1PlatformOperationsSmtpTestRoute = Route{
	Path:   "/api/v1/platformoperations/smtp/test",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.PlatformConfiguration,
	},
}

const V1StoreJobs = "V1StoreJobs"

var V1StoreJobsRoute = Route{
	Path:     "/api/v1/storejobs",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.StoreSyncManager,
	},
}
