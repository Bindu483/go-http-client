package charter

import (
	v1apps "k8s.io/api/apps/v1"
	v1beta1apps "k8s.io/api/apps/v1beta1"
	v1beta2apps "k8s.io/api/apps/v1beta2"
	v1autoscaling "k8s.io/api/autoscaling/v1"
	v2beta1autoscaling "k8s.io/api/autoscaling/v2beta1"
	v2beta2autoscaling "k8s.io/api/autoscaling/v2beta2"
	v1batch "k8s.io/api/batch/v1"
	v1beta1batch "k8s.io/api/batch/v1beta1"
	v2alpha1batch "k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	v1core "k8s.io/api/core/v1"
	v1beta1extensions "k8s.io/api/extensions/v1beta1"
	v1rbac "k8s.io/api/rbac/v1"
	v1alpha1rbac "k8s.io/api/rbac/v1alpha1"
	v1beta1rbac "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (c *RcHelmChart) RemoveUnsupportedManifests() {
	for k, k8sObjects := range c.DeploymentManifest.Chart {
		var t []K8sObject
		for _, obj := range k8sObjects {
			if isSupported := isObjectSupported(obj); isSupported {
				t = append(t, obj)
			}
		}
		if len(t) == 0 {
			c.DeploymentManifest.Chart[k] = nil
		} else {
			c.DeploymentManifest.Chart[k] = t
		}
	}
}

func isObjectSupported(object runtime.Object) bool {
	switch object.(type) {
	case *v1beta1batch.CronJob,
		*v2alpha1batch.CronJob,
		*v1batch.Job,
		*v1core.Pod,
		*v1apps.StatefulSet,
		*v1beta1apps.StatefulSet,
		*v1beta2apps.StatefulSet,
		*v1apps.Deployment,
		*v1beta1extensions.Deployment,
		*v1beta1apps.Deployment,
		*v1beta2apps.Deployment,
		*v1core.ServiceAccount,
		*v1rbac.Role,
		*v1beta1rbac.Role,
		*v1alpha1rbac.Role,
		*v1rbac.RoleBinding,
		*v1beta1rbac.RoleBinding,
		*v1alpha1rbac.RoleBinding,
		*v1core.ResourceQuota,
		*v1core.PersistentVolume,
		*v1autoscaling.HorizontalPodAutoscaler,
		*v2beta2autoscaling.HorizontalPodAutoscaler,
		*v2beta1autoscaling.HorizontalPodAutoscaler,
		*v1core.PersistentVolumeClaim,
		*v1core.Secret,
		*v1core.ConfigMap,
		*v1core.Service:
		return true
	}
	return false
}

func (c *RcHelmChart) RemoveResourceQuota() {
	for k, k8sObjects := range c.DeploymentManifest.Chart {
		var t []K8sObject
		for _, obj := range k8sObjects {
			switch obj.(type) {
			case *corev1.ResourceQuota:
				{
					break
				}
			default:
				t = append(t, obj)
			}
			if len(t) == 0 {
				c.DeploymentManifest.Chart[k] = nil
			} else {
				c.DeploymentManifest.Chart[k] = t
			}
		}
	}
}
