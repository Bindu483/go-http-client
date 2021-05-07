package models

type Resources struct {
	Cpu    *Resource `json:"cpu"`
	Memory *Resource `json:"memory"`
	Disk   *Resource `json:"disk"`
}

type Resource struct {
	Value int64 `json:"value"`
	Scale int32 `json:"scale"`
}
