package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1Health = "V1Health"

var V1HealthRoute = Route{
	Path:     "/health",
	Method:   http.MethodGet,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Public,
	},
}
