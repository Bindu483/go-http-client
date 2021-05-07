package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const PlatformUI = "PlatformUI"

var PlatformUIRoute = Route{
	Path:   "/",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Public,
	},
}
