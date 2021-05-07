package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

//TODO: this should ideally not be public, we made it public because we did not want to choke UMS for click stream eventing
const V1DataCollector = "V1DataCollector"

var V1DataCollectorRoute = Route{
	Path:   "/api/v1/data-collector",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V1DataCollectorCollector = "V1DataCollectorCollector"

var V1DataCollectorCollectorRoute = Route{
	Path:   "/api/v1/data-collector/collector",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Public,
	},
}
