package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1GetDeployment = "V1GetDeployment"

var V1GetDeploymentRoute = Route{
	Path:   "/api/v1/getDeployment",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1DeploymentListSolutions = "V1DeploymentListSolutions"

var V1DeploymentListSolutionsRoute = Route{
	Path:   "/api/v1/deployment/listSolutions",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}

const V1ListDeploymentsInCluster = "V1ListDeploymentsInCluster"

var V1ListDeploymentsInClusterRoute = Route{
	Path:   "/api/v1/listDeploymentsInCluster",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1DeploymentRequestRelease = "V1DeploymentRequestRelease"

var V1DeploymentRequestReleaseRoute = Route{
	Path:   "/api/v1/deployment/requestRelease",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}

const V1DeploymentSyncSolutions = "V1DeploymentSyncSolutions"

var V1DeploymentSyncSolutionsRoute = Route{
	Path:   "/api/v1/deployment/syncSolutions",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
	EnableTracing:    true,
	TracingEventName: "SyncApps",
}

const V1DeploymentGetAppDownloadStatus = "V1DeploymentGetAppDownloadStatus"

var V1DeploymentGetAppDownloadStatusRoute = Route{
	Path:   "/api/v1/deployment/getAppDownloadStatus",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.AppVersionManagement,
	},
}

const V1DeploymentGetAppVersionDetails = "V1DeploymentGetAppVersionDetails"

var V1DeploymentGetAppVersionDetailsRoute = Route{
	Path:   "/api/v1/deployment/getAppVersionDetails",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1ListDeployments = "V1ListDeployments"

var V1ListDeploymentsRoute = Route{
	Path:   "/api/v1/listDeployments",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1UpdateDeploymentConfig = "V1UpdateDeploymentConfig"

var V1UpdateDeploymentConfigRoute = Route{
	Path:   "/api/v1/updateDeploymentConfig",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1GetQosIndicatorForDeployment = "V1GetQosIndicatorForDeployment"

var V1GetQosIndicatorForDeploymentRoute = Route{
	Path:   "/api/v1/getQosIndicatorForDeployment",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1ListQosIndicatorForDeployment = "V1ListQosIndicatorForDeployment"

var V1ListQosIndicatorForDeploymentRoute = Route{
	Path:   "/api/v1/listQosIndicatorForDeployment",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1RemoveDeployment = "V1RemoveDeployment"

var V1RemoveDeploymentRoute = Route{
	Path:   "/api/v1/removeDeployment",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1ListDeploymentsOfApp = "V1ListDeploymentsOfApp"

var V1ListDeploymentsOfAppRoute = Route{
	Path:   "/api/v1/listDeploymentsofApp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V1IsAppDeployed = "V1IsAppDeployed"

var V1IsAppDeployedRoute = Route{
	Path:   "/api/v1/isAppDeployed",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1UpgradeApp = "V1UpgradeApp"

var V1UpgradeAppRoute = Route{
	Path:   "/api/v1/upgradeApp",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
}

const V1GetCompatibleVersions = "v1GetCompatibleVersions"

var V1GetCompatibleVersionsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodPost,
	Path:   "/api/v1/deployment/getCompatibleVersions",
}

const V1ListBackups = "v1ListBackups"

var V1ListBackupsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodPost,
	Path:   "/api/v1/deployment/listbackup",
}

const V1DeleteBackups = "v1DeleteBackups"

var V1DeleteBackupsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodDelete,
	Path:   "/api/v1/deployment/deletebackup",
}
