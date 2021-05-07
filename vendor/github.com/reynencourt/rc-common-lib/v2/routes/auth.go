package routes

import (
	"net/http"

	"github.com/reynencourt/rc-common-lib/v2/routes/permission"
)

const V2AuthServer = "V2AuthServer"

//TODO: how better to handle path prefixes

var V2AuthServerRoute = Route{
	Path:     "/api/v2/auth_server",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2Onboard = "V2Onboard"

var V2OnboardRoute = Route{
	Path:     "/api/v2/onboard",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V1UserProfile = "V1UserProfile"

var V1UserProfileRoute = Route{
	Path:     "/api/v1/user/profile",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V1UserInfo = "V1UserInfo"

var V1UserInfoRoute = Route{
	Path:     "/api/v1/user/info",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V2AuthServerLogin = "V2AuthServerLogin"

var V2AuthServerLoginRoute = Route{
	Path:   "/api/v2/auth_server/login",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2AuthServerLogout = "V2AuthServerLogout"

var V2AuthServerLogoutRoute = Route{
	Path:   "/api/v2/auth_server/logout",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2AuthServerAuthorize = "V2AuthServerAuthorize"

var V2AuthServerAuthorizeRoute = Route{
	Path:   "/api/v2/auth_server/authorize",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2AuthServerToken = "V2AuthServerToken"

var V2AuthServerTokenRoute = Route{
	Path:   "/api/v2/auth_server/token",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2AuthServerUserInfo = "V2AuthServerUserInfo"

var V2AuthServerUserInfoRoute = Route{
	Path:   "/api/v2/auth_server/userinfo",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2AuthServerRefresh = "V2AuthServerRefresh"

var V2AuthServerRefreshRoute = Route{
	Path:   "/api/v2/auth_server/refresh",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2AuthServerProviders = "V2AuthServerProviders"

var V2AuthServerProvidersRoute = Route{
	Path:   "/api/v2/auth_server/providers",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2AuthServerHealth = "V2AuthServerHealth"

var V2AuthServerHealthRoute = Route{
	Path:   "/api/v2/auth_server/health",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Public,
	},
}

const V2AuthServerChangePassword = "V2AuthServerChangePassword"

var V2AuthServerChangePasswordRoute = Route{
	Path:   "/api/v2/auth_server/change_password",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V2AuthServerChangePasswordProfile = "V2AuthServerChangePasswordProfile"

var V2AuthServerChangePasswordProfileRoute = Route{
	Path:   "/api/v2/auth_server/change_password_profile",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Basic,
	},
}

const V2AuthServerUsers = "V2AuthServerUsers"

var V2AuthServerUsersRoute = Route{
	Path:   "/api/v2/auth_server/users",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AuthServerUser = "V2AuthServerUser"

var V2AuthServerUserRoute = Route{
	Path:    "/api/v2/auth_server/user",
	Methods: AllMethods,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V2AuthServerDeactivateUser = "V2AuthServerDeactivateUser"

var V2AuthServerDeactivateUserRoute = Route{
	Path:   "/api/v2/auth_server/deactivate/user",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}

const V1SsoSamlLogin = "V1SsoSamlLogin"

var V1SsoSamlLoginRoute = Route{
	Path:   "/api/v1/sso/saml/login",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V1SsoSamlInitiate = "V1SsoSamlInitiate"

var V1SsoSamlInitiateRoute = Route{
	Path:   "/api/v1/sso/saml/initiate",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V1Sso = "V1Sso"

var V1SsoRoute = Route{
	Path:     "/api/v1/sso",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const SsoLogin = "SsoLogin"

var SsoLoginRoute = Route{
	Path:   "/sso/login",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2OauthLogin = "OauthLogin"

var V2OauthLoginRoute = Route{
	Path:   "/api/v2/auth_server/authorize",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2OauthToken = "OauthToken"

var V2OauthTokenRoute = Route{
	Path:   "/api/v2/auth_server/token",
	Method: http.MethodPost,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V2OauthUserInfo = "OauthUserInfo"

var V2OauthUserInfoRoute = Route{
	Path:   "/api/v2/auth_server/userinfo",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.Login,
	},
}

const V1SsoSettings = "V1SsoSettings"

var V1SsoSettingsRoute = Route{
	Path:   "/api/v1/sso/settings",
	Method: http.MethodGet,
	Permissions: []permission.Permission{
		permission.SSOConfiguration,
	},
}

const V1Permission = "V1Permission"

var V1PermissionRoute = Route{
	Path:     "/api/v1/permission",
	Methods:  AllMethods,
	IsPrefix: true,
	Permissions: []permission.Permission{
		permission.AccessControl,
	},
}
