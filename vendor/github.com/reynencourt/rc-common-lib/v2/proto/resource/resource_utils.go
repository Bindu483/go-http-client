package resource

import (
	"github.com/reynencourt/rc-common-lib/v2/commons"
	"k8s.io/apimachinery/pkg/api/resource"
	k8sResource "k8s.io/apimachinery/pkg/api/resource"
)

func (r *Resources) Add(y *Resources) {
	var yCPUMinAbs, yCPUMaxAbs, yMemMinAbs, yMemMaxAbs *resource.Quantity

	yCPUMinAbs = k8sResource.NewScaledQuantity(y.Cpu.Minimum, resource.Scale(y.Cpu.Scale))
	yCPUMaxAbs = k8sResource.NewScaledQuantity(y.Cpu.Maximum, resource.Scale(y.Cpu.Scale))
	yMemMinAbs = k8sResource.NewScaledQuantity(y.Memory.Minimum, resource.Scale(y.Memory.Scale))
	yMemMaxAbs = k8sResource.NewScaledQuantity(y.Memory.Maximum, resource.Scale(y.Memory.Scale))

	r.Memory.Minimum += yMemMinAbs.ScaledValue(resource.Scale(r.Memory.Scale))
	r.Memory.Maximum += yMemMaxAbs.ScaledValue(resource.Scale(r.Memory.Scale))
	r.Cpu.Maximum += yCPUMaxAbs.ScaledValue(resource.Scale(r.Cpu.Scale))
	r.Cpu.Minimum += yCPUMinAbs.ScaledValue(resource.Scale(r.Cpu.Scale))
}

func (r *Resources) Sub(y *Resources) {
	var yCPUMinAbs, yCPUMaxAbs, yMemMinAbs, yMemMaxAbs *resource.Quantity

	yCPUMinAbs = k8sResource.NewScaledQuantity(y.Cpu.Minimum, resource.Scale(y.Cpu.Scale))
	yCPUMaxAbs = k8sResource.NewScaledQuantity(y.Cpu.Maximum, resource.Scale(y.Cpu.Scale))
	yMemMinAbs = k8sResource.NewScaledQuantity(y.Memory.Minimum, resource.Scale(y.Memory.Scale))
	yMemMaxAbs = k8sResource.NewScaledQuantity(y.Memory.Maximum, resource.Scale(y.Memory.Scale))

	r.Memory.Minimum -= yMemMinAbs.ScaledValue(resource.Scale(r.Memory.Scale))
	r.Memory.Maximum -= yMemMaxAbs.ScaledValue(resource.Scale(r.Memory.Scale))
	r.Cpu.Maximum -= yCPUMaxAbs.ScaledValue(resource.Scale(r.Cpu.Scale))
	r.Cpu.Minimum -= yCPUMinAbs.ScaledValue(resource.Scale(r.Cpu.Scale))
}

func (r *NodeResource) Add(y *NodeResource) {
	yCpu := k8sResource.NewScaledQuantity(y.Cpu.Value, resource.Scale(y.Cpu.Scale))
	yMem := k8sResource.NewScaledQuantity(y.Memory.Value, resource.Scale(y.Memory.Scale))

	r.Cpu.Value += yCpu.ScaledValue(resource.Scale(r.Cpu.Scale))
	r.Memory.Value += yMem.ScaledValue(resource.Scale(r.Memory.Scale))
}

func (r *NodeResource) Sub(y *NodeResource) {
	yCpu := k8sResource.NewScaledQuantity(y.Cpu.Value, resource.Scale(y.Cpu.Scale))
	yMem := k8sResource.NewScaledQuantity(y.Memory.Value, resource.Scale(y.Memory.Scale))

	r.Cpu.Value -= yCpu.ScaledValue(resource.Scale(r.Cpu.Scale))
	r.Memory.Value -= yMem.ScaledValue(resource.Scale(r.Memory.Scale))
}

func (r *ClusterResourceInfo) RemoveResourceFromClusterResource(nodes []string) {
	for i, node := range r.Nodes.Node {
		if commons.Contains(nodes, node.Name) {
			r.Capacity.Sub(node.Capacity)
			r.Allocatable.Sub(node.Allocatable)
			r.Nodes.Node = append(r.Nodes.Node[0:i], r.Nodes.Node[i+1:]...)
		}
	}
}

func (r *ClusterResourceInfo) GetAvailableResources() *Resources {
	return &Resources{
		Cpu: &ResourceDefinition{
			Minimum: r.Allocatable.Cpu.Value - r.Used.Cpu.Maximum,
			Maximum: r.Allocatable.Cpu.Value - r.Used.Cpu.Minimum,
			Scale:   int32(k8sResource.Milli),
		},
		Memory: &ResourceDefinition{
			Minimum: r.Allocatable.Memory.Value - r.Used.Memory.Maximum,
			Maximum: r.Allocatable.Memory.Value - r.Used.Memory.Minimum,
			Scale:   int32(k8sResource.Mega),
		},
	}
}
