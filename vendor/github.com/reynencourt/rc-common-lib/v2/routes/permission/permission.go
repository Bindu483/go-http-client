package permission

type Permission string

// Auth modules constants
const (
	AppVersionManagement  = "app_version_management"
	ClusterManagement     = "cluster_management"
	AccessControl         = "access_control"
	Public                = "public"
	SSOConfiguration      = "sso_configuration"
	Login                 = "login"
	BroadcastManagement   = "broadcast_management"
	ViewStoreSolutions    = "view_store_solutions"
	StoreSyncManager      = "store_sync_manager"
	ProjectViewer         = "project_viewer"
	ProjectAdmin          = "project_admin"
	ProjectManagement     = "project_management"
	PlatformConfiguration = "platform_configuration"
	Basic                 = "basic"
	TestDriveAdmin        = "test_drive_admin"
	StorageManagement     = "storage_management"
)
