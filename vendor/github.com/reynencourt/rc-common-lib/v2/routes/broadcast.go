package routes

import (
	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V1Broadcast = "V1Broadcast"

var V1BroadcastRoute = Route{
	Path:     "/api/v1/broadcast",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.BroadcastManagement,
	},
}
