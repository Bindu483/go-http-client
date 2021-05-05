package charter

import (
	"github.com/pkg/errors"
	rcResource "github.com/reynencourt/rc-common-lib/v2/proto/resource"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (c *RcHelmChart) CalculateContention(cr *rcResource.ClusterResourceInfo) (*rcResource.DeploymentQosIndicator, error) {
	rq, err := c.ResourceQuotaExist()
	if err != nil {
		return nil, err
	}
	if rq == nil {
		return nil, errors.New("failed to get app resource info")
	}
	cpuLimit := rq.Spec.Hard[LimitCPU]
	cpuMin := rq.Spec.Hard[RequestCPU]
	memLim := rq.Spec.Hard[LimitMemory]
	memReq := rq.Spec.Hard[RequestMemory]

	requiredCpuMax := float64(cpuLimit.ScaledValue(resource.Scale(cr.Allocatable.Cpu.Scale)) + cr.Allocatable.Cpu.Value)
	requiredCpuMin := float64(cpuMin.ScaledValue(resource.Scale(cr.Allocatable.Cpu.Scale)) + cr.Used.Cpu.Minimum)
	requiredMemMax := float64(memLim.ScaledValue(resource.Scale(cr.Allocatable.Memory.Scale)) + cr.Used.Memory.Maximum)
	requiredMemMin := float64(memReq.ScaledValue(resource.Scale(cr.Allocatable.Memory.Scale)) + cr.Used.Memory.Minimum)

	totalAllocatableCpu := float64(cr.Allocatable.Cpu.Value)
	totalAllocatableMem := float64(cr.Allocatable.Memory.Value)

	var cpuQosIndicator, memQosIndicator float64
	var canDeploy, hasSufficientCpu, hasSufficientMem = true, true, true

	if requiredCpuMin == requiredCpuMax {
		cpuQosIndicator = rcResource.Sgn(requiredCpuMax - totalAllocatableCpu)
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
		memQosIndicator = rcResource.Sgn(requiredMemMax - totalAllocatableMem)
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

	return &rcResource.DeploymentQosIndicator{
		CanDeploy:       canDeploy,
		CpuQosIndicator: cpuQosIndicator,
		MemQosIndicator: memQosIndicator,
	}, nil
}
