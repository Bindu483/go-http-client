package routes

import (
	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
	"net/http"
)

const V2ListBackups = "V2ListBackups"

var V2ListBackupsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodGet,
	Path:   "/api/v2/deployments/{deployment_id}/backup",
}

const V2DeleteBackups = "V2DeleteBackups"

var V2DeleteBackupsRoute = Route{
	Permissions: []permission.Permission{
		permission.ProjectAdmin,
	},
	Method: http.MethodDelete,
	Path:   "/api/v2/deployments/{deployment_id}/backup/{backup_name}",
}
