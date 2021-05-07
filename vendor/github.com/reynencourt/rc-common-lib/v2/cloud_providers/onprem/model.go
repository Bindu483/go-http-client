package onprem

import "github.com/reynencourt/rc-common-lib/v2/cloud_providers"

type Provider struct {
	Input  Input
	Output Output
}

type Input struct {
	LocalInstallation bool                   `json:"local_installation" yaml:"local_installation"`
	MachineHostOrIP   string                 `json:"machine_ips" yaml:"machine_ips"`
	MachineUserName   string                 `json:"machine_user" yaml:"machine_user"`
	PrivateKeyName    string                 `json:"private_key_path" yaml:"private_key_path"`
	OsType            cloud_providers.OsType `json:"operating_system" yaml:"operating_system"`
}

type Output struct {
}
