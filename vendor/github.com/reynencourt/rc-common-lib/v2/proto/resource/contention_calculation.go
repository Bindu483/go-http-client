package resource

import "github.com/reynencourt/rc-common-lib/v2/commons"

const (
	DefaultCpuLimit       = "1000m"
	DefaultMemoryLimit    = "1000M"
	DefaultCpuRequest     = "500m"
	DefaultMemeoryRequest = "500M"
)

func (r *ClusterResourceInfo) WillCauseContention(nodes []string) *NodeDeletionQosIndicator {
	var totalCpuToBeRemoved, totalMemToBeRemoved int64

	for _, n := range r.Nodes.Node {
		if commons.Contains(nodes, n.Name) {
			totalCpuToBeRemoved += n.Allocatable.Cpu.Value
			totalMemToBeRemoved += n.Allocatable.Memory.Value
		}
	}
	totalUsedCpuMax := float64(r.Used.Cpu.Maximum)
	totalUsedCpuMin := float64(r.Used.Cpu.Minimum)
	totalUsedMemMax := float64(r.Used.Memory.Maximum)
	totalUsedMemMin := float64(r.Used.Memory.Minimum)
	totalAllocatableCpu := float64(r.Allocatable.Cpu.Value)
	totalAllocatableMem := float64(r.Allocatable.Memory.Value)

	totalCpuToBeRemovedFloat := float64(totalCpuToBeRemoved)
	totalMemToBeRemovedFloat := float64(totalMemToBeRemoved)

	var cpuQosIndicator, memQosIndicator float64
	var canDelete, hasSufficientCpu, hasSufficientMem = true, true, true

	if totalUsedCpuMax == totalUsedCpuMin {
		cpuQosIndicator = Sgn(totalUsedCpuMax - (totalAllocatableCpu - totalCpuToBeRemovedFloat))
		if cpuQosIndicator > 0 {
			hasSufficientCpu = false
		}
	} else {
		cpuQosIndicator := (totalUsedCpuMax - (totalAllocatableCpu - totalCpuToBeRemovedFloat)) /
			(totalUsedCpuMax - totalUsedCpuMin)
		if cpuQosIndicator > 1 {
			hasSufficientCpu = false
		}
	}

	if totalUsedMemMax == totalUsedMemMin {
		memQosIndicator = Sgn(totalUsedMemMax - (totalAllocatableMem - totalMemToBeRemovedFloat))
		if memQosIndicator > 0 {
			hasSufficientMem = false
		}
	} else {
		memQosIndicator := (totalUsedMemMax - (totalAllocatableMem - totalMemToBeRemovedFloat)) /
			(totalUsedMemMax - totalUsedMemMin)
		if memQosIndicator > 1 {
			hasSufficientMem = false
		}
	}

	if !hasSufficientCpu || !hasSufficientMem {
		canDelete = false
	}

	return &NodeDeletionQosIndicator{
		CanDelete:       canDelete,
		CpuQosIndicator: cpuQosIndicator,
		MemQosIndicator: memQosIndicator,
	}
}

//No scale explit checks has this is all maintained by internal lib assumption is scale is always same
//good to have explicit scale conversions
func (r *ClusterResourceInfo) CalculateContention(res *Resources) *DeploymentQosIndicator {
	requiredCpuMax := float64(res.Cpu.Maximum + r.Used.Cpu.Maximum)
	requiredCpuMin := float64(res.Cpu.Minimum + r.Used.Cpu.Minimum)
	requiredMemMax := float64(res.Memory.Maximum + r.Used.Memory.Maximum)
	requiredMemMin := float64(res.Memory.Minimum + r.Used.Memory.Minimum)

	totalAllocatableCpu := float64(r.Allocatable.Cpu.Value)
	totalAllocatableMem := float64(r.Allocatable.Memory.Value)

	var cpuQosIndicator, memQosIndicator float64
	var canDeploy, hasSufficientCpu, hasSufficientMem = true, true, true

	if requiredCpuMin == requiredCpuMax {
		cpuQosIndicator = Sgn(requiredCpuMax - totalAllocatableCpu)
		if cpuQosIndicator > 0 {
			hasSufficientCpu = false
		}
	} else {
		cpuQosIndicator = (requiredCpuMax - totalAllocatableCpu) / (requiredCpuMax - requiredCpuMin)
		if cpuQosIndicator > 1 {
			hasSufficientCpu = false
		}
	}

	if requiredMemMin == requiredMemMax {
		memQosIndicator = Sgn(requiredMemMax - totalAllocatableMem)
		if memQosIndicator > 0 {
			hasSufficientMem = false
		}
	} else {
		memQosIndicator = (requiredMemMax - totalAllocatableMem) / (requiredMemMax - requiredMemMin)
		if memQosIndicator > 1 {
			hasSufficientMem = false
		}
	}

	if !hasSufficientCpu || !hasSufficientMem {
		canDeploy = false
	}

	return &DeploymentQosIndicator{
		CanDeploy:       canDeploy,
		CpuQosIndicator: cpuQosIndicator,
		MemQosIndicator: memQosIndicator,
	}
}

func Sgn(a float64) float64 {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
