package charter

import (
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

func (c *RcHelmChart) ResourceQuotaExist() (*corev1.ResourceQuota, error) {
	for _, k8sObjects := range c.DeploymentManifest.Chart {
		for _, obj := range k8sObjects {
			switch obj.(type) {
			case *corev1.ResourceQuota:
				rq, ok := obj.(*corev1.ResourceQuota)
				if !ok {
					return nil, errors.New("could not parse resource quota definition")
				}
				if memLimit := rq.Spec.Hard[LimitMemory]; memLimit.IsZero() {
					return nil, errors.New("limits for memory are zero")
				}
				if cpuLimit := rq.Spec.Hard[LimitCPU]; cpuLimit.IsZero() {
					return nil, errors.New("limits for CPU are zero")
				}
				return rq, nil
			}
		}
	}
	return nil, ResourceRequirementsNotDefined
}
