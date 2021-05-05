package charter

import (
	"errors"
	"math"

	"github.com/reynencourt/rc-common-lib/v2/ops/kube"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const HeadSpace = 0.25

const (
	RequestCPU    = "requests.cpu"
	LimitCPU      = "limits.cpu"
	RequestMemory = "requests.memory"
	LimitMemory   = "limits.memory"
	Memory        = "memory"
	Cpu           = "cpu"
	Storage       = "storage"
)

var ResourceRequirementsNotDefined = errors.New("resource definiton not found")

func addExtraResource(rl map[string]*resource.Quantity, quantity float64) error {
	for k, v := range rl {
		if k == LimitCPU || k == RequestCPU {
			// I had to handle cases like 500m which is why I get scaled value and when i create new resource i divide by 1000
			// I tested with CPU=50000 and i did not see .MilliValue() overflowing
			milliValueFloat := float64(v.ScaledValue(resource.Milli))
			scaledValue := (milliValueFloat * 0.25) + milliValueFloat
			rl[k] = resource.NewScaledQuantity(resource.NewQuantity(int64(scaledValue), v.Format).Value(), resource.Milli)
		} else {
			// sdk has a funny way to parse the values to different scales.
			// if you parse 1.5M to Mega the output is 2
			// if you parse 1.25M to Mega the output is 1
			// if you parse 1.5K to Mega there is a panic
			// I am assuming the memory will always be in Megabytes so i scale the given value to Kilo so the final values becomes decimal
			// given the specified value is is upto 3 floating point accuracy.
			// with this approach the calculated values are close to the desired values (no Math.Ceil)
			kiloValue := float64(v.ScaledValue(resource.Kilo))
			scaledValue := (kiloValue * quantity) + kiloValue
			rl[k] = resource.NewScaledQuantity(
				resource.NewQuantity(int64(scaledValue), v.Format).Value()/
					int64(math.Pow10(int(resource.Kilo))), resource.Mega)
		}
	}
	return nil
}

func addResources(cs []corev1.Container, rl map[string]*resource.Quantity, replicas *int32) (err error) {
	var rep int64
	if replicas == nil {
		rep = 1
	} else {
		rep = int64(*replicas)
	}
	for _, c := range cs {
		//if limits are null or zero values it means there are two possibilities,
		// 1. they specify requests so asking minimum, but the app can grow to whatever is available
		// 2. they dont want to set any resource definitions
		// in both the case i appply defaults, so even though requests are set
		// but limits are left out those requests are ignored
		var resourceRequirements corev1.ResourceRequirements
		if lim := c.Resources.Limits; lim == nil {
			resourceRequirements = getDefaultResource()
		} else {
			resourceRequirements = c.Resources
		}

		defaultCpuLimit := resource.MustParse(kube.DefaultCpuLimit)
		defaultMemLimit := resource.MustParse(kube.DefaultMemoryLimit)
		defaultCpuRequest := resource.MustParse(kube.DefaultCpuRequest)
		defaultMemRequest := resource.MustParse(kube.DefaultMemoryRequest)
		limCpu := resourceRequirements.Limits[Cpu]
		limMem := resourceRequirements.Limits[Memory]

		if !limCpu.IsZero() {
			q := resource.NewQuantity(limCpu.ScaledValue(resource.Milli)*rep, limCpu.Format)
			rl[LimitCPU].Add(*resource.NewScaledQuantity(q.Value(), resource.Milli))
		} else {
			q := resource.NewQuantity(defaultCpuLimit.ScaledValue(resource.Milli)*rep, resource.DecimalSI)
			rl[LimitCPU].Add(*resource.NewScaledQuantity(q.Value(), resource.Milli))
		}
		if !limMem.IsZero() {
			q := resource.NewQuantity(limMem.ScaledValue(resource.Mega)*rep, limMem.Format)
			rl[LimitMemory].Add(*resource.NewScaledQuantity(q.Value(), resource.Mega))
		} else {
			q := resource.NewQuantity(defaultMemLimit.ScaledValue(resource.Mega)*rep, resource.BinarySI)
			rl[LimitMemory].Add(*resource.NewScaledQuantity(q.Value(), resource.Mega))
		}

		var reqCpu, reqMem resource.Quantity

		if resourceRequirements.Requests == nil {
			reqCpu = resourceRequirements.Limits[Cpu]
			reqMem = resourceRequirements.Limits[Memory]
		} else {
			reqCpu = resourceRequirements.Requests[Cpu]
			reqMem = resourceRequirements.Requests[Memory]
		}

		if !reqMem.IsZero() {
			q := resource.NewQuantity(reqMem.ScaledValue(resource.Mega)*rep, reqMem.Format)
			rl[RequestMemory].Add(*resource.NewScaledQuantity(q.Value(), resource.Mega))
		} else {
			q := resource.NewQuantity(defaultMemRequest.ScaledValue(resource.Mega)*rep, resource.BinarySI)
			rl[RequestMemory].Add(*resource.NewScaledQuantity(q.Value(), resource.Mega))
		}

		if !reqCpu.IsZero() {
			q := resource.NewQuantity(reqCpu.ScaledValue(resource.Milli)*rep, reqCpu.Format)
			rl[RequestCPU].Add(*resource.NewScaledQuantity(q.Value(), resource.Milli))
		} else {
			q := resource.NewQuantity(defaultCpuRequest.ScaledValue(resource.Milli)*rep, resource.DecimalSI)
			rl[RequestCPU].Add(*resource.NewScaledQuantity(q.Value(), resource.Milli))
		}
	}

	return nil
}

func getDefaultResource() corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			LimitMemory: resource.MustParse(kube.DefaultMemoryLimit),
			LimitCPU:    resource.MustParse(kube.DefaultCpuLimit),
		},
		Requests: corev1.ResourceList{
			RequestMemory: resource.MustParse(kube.DefaultMemoryRequest),
			RequestCPU:    resource.MustParse(kube.DefaultCpuRequest),
		},
	}
}

func addResourcesStrict(cs []corev1.Container, rl map[string]*resource.Quantity, replicas *int32) (err error) {
	var rep int64
	if replicas == nil {
		rep = 1
	} else {
		rep = int64(*replicas)
	}
	for _, c := range cs {
		if lim := c.Resources.Limits; lim != nil {
			if !lim.Cpu().IsZero() {
				q := resource.NewQuantity(lim.Cpu().AsDec().UnscaledBig().Int64()*rep, lim.Cpu().Format)
				rl[LimitCPU].Add(*q)
			} else {
				return ResourceRequirementsNotDefined
			}
			if !lim.Memory().IsZero() {
				q := resource.NewQuantity(lim.Memory().Value()*rep, lim.Memory().Format)
				rl[LimitMemory].Add(*q)
			} else {
				return ResourceRequirementsNotDefined
			}
		} else {
			return ResourceRequirementsNotDefined
		}

		req := c.Resources.Requests
		if req == nil {
			req = c.Resources.Limits
		}
		if !req.Memory().IsZero() {
			q := resource.NewQuantity(req.Memory().Value()*rep, req.Memory().Format)
			rl[RequestMemory].Add(*q)
		} else {
			return ResourceRequirementsNotDefined
		}
		if !req.Cpu().IsZero() {
			q := resource.NewQuantity(req.Cpu().AsDec().UnscaledBig().Int64()*rep, req.Cpu().Format)
			rl[RequestCPU].Add(*q)
		} else {
			return ResourceRequirementsNotDefined
		}
	}
	return nil
}
