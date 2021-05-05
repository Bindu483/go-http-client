package deploymentManager

import (
	"github.com/reynencourt/rc-common-lib/v2/app/charter"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (r *ResourceInfo) ToV1ResourceQuotaSpec() v1.ResourceQuotaSpec {
	return v1.ResourceQuotaSpec{
		Hard: v1.ResourceList{
			charter.LimitCPU:      resource.NewScaledQuantity(r.Requested.Cpu.Maximum, resource.Scale(r.Requested.Cpu.Scale)).DeepCopy(),
			charter.RequestCPU:    resource.NewScaledQuantity(r.Requested.Cpu.Maximum, resource.Scale(r.Requested.Cpu.Scale)).DeepCopy(),
			charter.LimitMemory:   resource.NewScaledQuantity(r.Requested.Memory.Maximum, resource.Scale(r.Requested.Memory.Scale)).DeepCopy(),
			charter.RequestMemory: resource.NewScaledQuantity(r.Requested.Memory.Minimum, resource.Scale(r.Requested.Memory.Scale)).DeepCopy(),
		},
	}
}
