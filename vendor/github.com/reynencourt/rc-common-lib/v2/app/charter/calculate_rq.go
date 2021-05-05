package charter

import (
	"github.com/pkg/errors"
	v1apps "k8s.io/api/apps/v1"
	v1beta1apps "k8s.io/api/apps/v1beta1"
	v1beta2apps "k8s.io/api/apps/v1beta2"
	v1batch "k8s.io/api/batch/v1"
	v1beta1batch "k8s.io/api/batch/v1beta1"
	v2alpha1batch "k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	v1beta1extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

//If a Container specifies its own memory limit, but does not specify a memory request, Kubernetes
//automatically assigns a memory request that matches the limit. Similarly, if a Container specifies
//its own CPU limit, but does not specify a CPU request, Kubernetes automatically assigns a CPU request that matches the limit.

const (
	RcResourceQuotaName = "resource-quota"
	OldValue            = "old_value"
)

func (c RcHelmChart) CalculateResourceQuota() (*corev1.ResourceQuota, error) {
	var resourceList = map[string]*resource.Quantity{
		RequestCPU:    resource.NewQuantity(0, resource.DecimalSI),
		LimitCPU:      resource.NewQuantity(0, resource.DecimalSI),
		RequestMemory: resource.NewQuantity(0, resource.BinarySI),
		LimitMemory:   resource.NewQuantity(0, resource.BinarySI),
	}
	var err error
	for _, k8sObjects := range c.DeploymentManifest.Chart {
		for _, obj := range k8sObjects {
			switch obj.(type) {
			case *corev1.Pod:
				p, ok := obj.(*corev1.Pod)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Containers, resourceList, nil)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta1extensions.ReplicaSet:
				p, ok := obj.(*v1beta1extensions.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1apps.ReplicaSet:
				p, ok := obj.(*v1apps.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta2apps.ReplicaSet:
				p, ok := obj.(*v1beta2apps.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *corev1.ReplicationController:
				p, ok := obj.(*corev1.ReplicationController)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1apps.Deployment:
				p, ok := obj.(*v1apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta2apps.Deployment:
				p, ok := obj.(*v1beta2apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta1apps.Deployment:
				p, ok := obj.(*v1beta1apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta1extensions.Deployment:
				p, ok := obj.(*v1beta1extensions.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1apps.StatefulSet:
				p, ok := obj.(*v1apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta2apps.StatefulSet:
				p, ok := obj.(*v1beta2apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta1apps.StatefulSet:
				p, ok := obj.(*v1beta1apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Replicas)
				if err != nil {
					return nil, err
				}
				break
			case *v1apps.DaemonSet:
				return nil, errors.New("daemonsets are not allowed")
			case *v1beta1extensions.DaemonSet:
				return nil, errors.New("daemonsets are not allowed")
			case *v1beta2apps.DaemonSet:
				return nil, errors.New("daemonsets are not allowed")
			case *v1batch.Job:
				p, ok := obj.(*v1batch.Job)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.Template.Spec.Containers, resourceList, p.Spec.Parallelism)
				if err != nil {
					return nil, err
				}
				break
			case *v1beta1batch.CronJob:
				p, ok := obj.(*v1beta1batch.CronJob)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.JobTemplate.Spec.Template.Spec.Containers, resourceList, p.Spec.JobTemplate.Spec.Parallelism)
				if err != nil {
					return nil, err
				}
				break
			case *v2alpha1batch.CronJob:
				p, ok := obj.(*v2alpha1batch.CronJob)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				err = addResources(p.Spec.JobTemplate.Spec.Template.Spec.Containers, resourceList, p.Spec.JobTemplate.Spec.Parallelism)
				if err != nil {
					return nil, err
				}
				break
			}

		}
	}
	err = addExtraResource(resourceList, HeadSpace)
	if err != nil {
		return nil, err
	}

	resourceQuota := &corev1.ResourceQuota{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ResourceQuota",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: RcResourceQuotaName,
			Labels: map[string]string{
				"created-by": "rc-user",
			},
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: corev1.ResourceList{
				LimitMemory:   resourceList[LimitMemory].DeepCopy(),
				LimitCPU:      resourceList[LimitCPU].DeepCopy(),
				RequestMemory: resourceList[RequestMemory].DeepCopy(),
				RequestCPU:    resourceList[RequestCPU].DeepCopy(),
			},
		},
	}
	return resourceQuota, nil
}

//Adds resources y to x and and x is written as annotations in the returned values
func AddResourceQuotas(x, y *corev1.ResourceQuota) (*corev1.ResourceQuota, error) {
	xLimitMem := x.Spec.Hard[LimitMemory]
	xRequestMem := x.Spec.Hard[RequestMemory]
	xLimitCpu := x.Spec.Hard[LimitCPU]
	xRequestCpu := x.Spec.Hard[RequestCPU]

	oldValue, err := json.Marshal(x.Spec.Hard)
	if err != nil {
		return nil, errors.New("failed to marshall resource_quota")
	}

	if len(x.Annotations) == 0 {
		x.Annotations = make(map[string]string)
	}
	x.Annotations[OldValue] = string(oldValue)

	xLimitCpu.Add(y.Spec.Hard[LimitCPU])
	xLimitMem.Add(y.Spec.Hard[LimitMemory])
	xRequestCpu.Add(y.Spec.Hard[RequestCPU])
	xRequestMem.Add(y.Spec.Hard[RequestMemory])

	x.Spec.Hard[LimitMemory] = xLimitMem
	x.Spec.Hard[LimitCPU] = xLimitCpu
	x.Spec.Hard[RequestMemory] = xRequestMem
	x.Spec.Hard[RequestCPU] = xRequestCpu

	return x, nil
}

func SubtractResourceQuotas(x, y *corev1.ResourceQuota) (*corev1.ResourceQuota, error) {
	xLimitMem := x.Spec.Hard[LimitMemory]
	xRequestMem := x.Spec.Hard[RequestMemory]
	xLimitCpu := x.Spec.Hard[LimitCPU]
	xRequestCpu := x.Spec.Hard[RequestCPU]

	oldValue, err := json.Marshal(x.Spec.Hard)
	if err != nil {
		return nil, errors.New("failed to marshall resource_quota")
	}

	if len(x.Annotations) == 0 {
		x.Annotations = make(map[string]string)
	}
	x.Annotations[OldValue] = string(oldValue)

	xLimitCpu.Sub(y.Spec.Hard[LimitCPU])
	xLimitMem.Sub(y.Spec.Hard[LimitMemory])
	xRequestCpu.Sub(y.Spec.Hard[RequestCPU])
	xRequestMem.Sub(y.Spec.Hard[RequestMemory])

	x.Spec.Hard[LimitMemory] = xLimitMem
	x.Spec.Hard[LimitCPU] = xLimitCpu
	x.Spec.Hard[RequestMemory] = xRequestMem
	x.Spec.Hard[RequestCPU] = xRequestCpu

	return x, nil
}
