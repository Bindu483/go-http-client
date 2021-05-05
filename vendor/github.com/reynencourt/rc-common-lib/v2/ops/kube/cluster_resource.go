package kube

import (
	rcResource "github.com/reynencourt/rc-common-lib/v2/proto/resource"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (nodes K8sNodes) CalculateClusterResource(cr *rcResource.ClusterResourceInfo) *rcResource.ClusterResourceInfo {
	var cpuCap, memCap, allCPU, allMem, allDisk, diskCap resource.Quantity
	for i, node := range nodes {

		cpuCap.Add(node.Capacity[v1.ResourceCPU])
		memCap.Add(node.Capacity[v1.ResourceMemory])
		allCPU.Add(node.Allocatable[v1.ResourceCPU])
		allMem.Add(node.Allocatable[v1.ResourceMemory])
		allDisk.Add(node.Allocatable[v1.ResourceEphemeralStorage])
		diskCap.Add(node.Capacity[v1.ResourceEphemeralStorage])

		cr.Nodes.Node[i] = &rcResource.NodeInfo{
			Name: node.Name,
			Ip:   node.InternalIP,
			Role: string(node.Role),
			Capacity: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: node.Capacity.Cpu().ToDec().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: node.Capacity.Memory().ToDec().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
				Disk: &rcResource.Resource{
					Value: node.Capacity.StorageEphemeral().ToDec().ScaledValue(resource.Giga),
					Scale: int32(resource.Giga),
				},
			},
			Allocatable: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: node.Allocatable.Cpu().ToDec().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: node.Allocatable.Memory().ToDec().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
				Disk: &rcResource.Resource{
					Value: node.Allocatable.StorageEphemeral().ToDec().ScaledValue(resource.Giga),
					Scale: int32(resource.Giga),
				},
			},
		}
	}
	cr.Capacity = &rcResource.NodeResource{
		Cpu: &rcResource.Resource{
			Value: cpuCap.ToDec().ScaledValue(resource.Milli),
			Scale: int32(resource.Milli),
		},
		Memory: &rcResource.Resource{
			Value: memCap.ToDec().ScaledValue(resource.Mega),
			Scale: int32(resource.Mega),
		},
		Disk: &rcResource.Resource{
			Value: diskCap.ToDec().ScaledValue(resource.Giga),
			Scale: int32(resource.Giga),
		},
	}
	cr.Allocatable = &rcResource.NodeResource{
		Cpu: &rcResource.Resource{
			Value: allCPU.ToDec().ScaledValue(resource.Milli),
			Scale: int32(resource.Milli),
		},
		Memory: &rcResource.Resource{
			Value: allMem.ToDec().ScaledValue(resource.Mega),
			Scale: int32(resource.Mega),
		},
		Disk: &rcResource.Resource{
			Value: allDisk.ToDec().ScaledValue(resource.Giga),
			Scale: int32(resource.Giga),
		},
	}
	zeroValue := resource.MustParse("0")
	cr.Used = &rcResource.Resources{
		Cpu: &rcResource.ResourceDefinition{
			Minimum: zeroValue.ScaledValue(resource.Milli),
			Maximum: zeroValue.ScaledValue(resource.Milli),
			Scale:   int32(resource.Milli),
		},
		Memory: &rcResource.ResourceDefinition{
			Minimum: zeroValue.ScaledValue(resource.Mega),
			Maximum: zeroValue.ScaledValue(resource.Mega),
			Scale:   int32(resource.Mega),
		},
	}
	return cr
}

func (nodes K8sNodes) AddResourceToClusterResource(crInfo *rcResource.ClusterResourceInfo) {
	for _, v := range nodes {
		node := &rcResource.NodeInfo{
			Name: v.Name,
			Ip:   v.InternalIP,
			Role: string(v.Role),
			Capacity: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: v.Capacity.Cpu().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: v.Capacity.Memory().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
			},
			Allocatable: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: v.Allocatable.Cpu().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: v.Allocatable.Memory().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
			},
		}
		crInfo.Nodes.Node = append(crInfo.Nodes.Node, node)
		crInfo.Capacity.Add(node.Capacity)
		crInfo.Allocatable.Add(node.Allocatable)
	}
}

func (nodes K8sNodes) RemoveResourceFromClusterResource(crInfo *rcResource.ClusterResourceInfo) {
	for i, v := range nodes {
		node := &rcResource.NodeInfo{
			Name: v.Name,
			Ip:   v.InternalIP,
			Role: string(v.Role),
			Capacity: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: v.Capacity.Cpu().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: v.Capacity.Memory().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
			},
			Allocatable: &rcResource.NodeResource{
				Cpu: &rcResource.Resource{
					Value: v.Allocatable.Cpu().ScaledValue(resource.Milli),
					Scale: int32(resource.Milli),
				},
				Memory: &rcResource.Resource{
					Value: v.Allocatable.Memory().ScaledValue(resource.Mega),
					Scale: int32(resource.Mega),
				},
			},
		}
		crInfo.Nodes.Node = append(crInfo.Nodes.Node[0:i], crInfo.Nodes.Node[i+1:]...)
		crInfo.Capacity.Sub(node.Capacity)
		crInfo.Allocatable.Sub(node.Allocatable)
	}
}
