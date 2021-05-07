package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2Abac = "V2Abac"

var V2AbacRoute = Route{
	Path:     "/api/v2/abac",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AbacPolicy = "V2AbacPolicy"

var V2AbacPolicyRoute = Route{
	Path:     "/api/v2/abac/policy",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AbacRole = "V2AbacRole"

var V2AbacRoleRoute = Route{
	Path:     "/api/v2/abac/role",
	Methods:  []string{http.MethodPut, http.MethodDelete, http.MethodPost},
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AbacRoles = "V2AbacRoles"

var V2AbacRolesRoute = Route{
	Path:   "/api/v2/abac/roles",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AbacResourceGroups = "V2AbacResourceGroups"

var V2AbacResourceGroupsRoute = Route{
	Path:    "/api/v2/abac/resource_groups",
	Methods: AllMethods,
	Permissions: []permission.Permission{
		permission.ProjectViewer,
	},
}

const V2AbacProjects = "V2AbacProjects"

var V2AbacProjectsRoute = Route{
	Path:    "/api/v2/abac/projects",
	Methods: AllMethods,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AbacManageProjects = "V2AbacManageProjects"

var V2AbacManageProjectsRoute = Route{
	Path:     "/api/v2/abac/manage/projects",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.ProjectManagement,
	},
}
