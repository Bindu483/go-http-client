package cloud_providers

type OsType string

const (
	Ubuntu1804 OsType = "ubuntu1804"
	CoreOS     OsType = "coreos"
	Ubuntu1604 OsType = "ubuntu1604"
	CentOS7    OsType = "centos7"
	CentOS8    OsType = "centos8"
	RHEL7      OsType = "rhel7"
	RHEL8      OsType = "rhel8"
	OSUnknown  OsType = "unknown"
)
