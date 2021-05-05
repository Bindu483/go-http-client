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
)

func (c *RcHelmChart) EnsureWorkloadsHasImagePullSecrets(secretName string) error {
	for _, k8sObjects := range c.DeploymentManifest.Chart {
		for _, obj := range k8sObjects {
			switch obj.(type) {
			case *corev1.Pod:
				p, ok := obj.(*corev1.Pod)
				if !ok {
					return errors.New("could not parse pod definition")
				}
				p.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta1extensions.ReplicaSet:
				r, ok := obj.(*v1beta1extensions.ReplicaSet)
				if !ok {
					return errors.New("could not parse replica set definition")
				}
				r.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1apps.ReplicaSet:
				r, ok := obj.(*v1apps.ReplicaSet)
				if !ok {
					return errors.New("could not parse replica set definition")
				}
				r.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
			case *v1beta2apps.ReplicaSet:
				r, ok := obj.(*v1beta2apps.ReplicaSet)
				if !ok {
					return errors.New("could not parse v1beta2apps replica set definition")
				}
				r.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
			case *corev1.ReplicationController:
				r, ok := obj.(*corev1.ReplicationController)
				if !ok {
					return errors.New("could not parse corev1 replica set definition")
				}
				r.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1apps.Deployment:
				d, ok := obj.(*v1apps.Deployment)
				if !ok {
					return errors.New("could not parse v1apps deployment definition")
				}
				d.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta2apps.Deployment:
				d, ok := obj.(*v1beta2apps.Deployment)
				if !ok {
					return errors.New("could not parse v1beta2apps deployment definition")
				}
				d.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta1apps.Deployment:
				d, ok := obj.(*v1beta1apps.Deployment)
				if !ok {
					return errors.New("could not parse v1beta1apps deployment definition")
				}
				d.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta1extensions.Deployment:
				d, ok := obj.(*v1beta1extensions.Deployment)
				if !ok {
					return errors.New("could not parse v1beta1extensions deployment definition")
				}
				d.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1apps.StatefulSet:
				s, ok := obj.(*v1apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1apps statefulset definition")
				}
				s.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta2apps.StatefulSet:
				s, ok := obj.(*v1beta2apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1beta2apps statefulset definition")
				}
				s.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta1apps.StatefulSet:
				s, ok := obj.(*v1beta1apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1beta1apps statefulset definition")
				}
				s.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1apps.DaemonSet:
				return errors.New("daemonsets are not allowed")
			case *v1beta1extensions.DaemonSet:
				return errors.New("daemonsets are not allowed")
			case *v1beta2apps.DaemonSet:
				return errors.New("daemonsets are not allowed")
			case *v1batch.Job:
				j, ok := obj.(*v1batch.Job)
				if !ok {
					return errors.New("could not parse v1batch job definition")
				}
				j.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v1beta1batch.CronJob:
				j, ok := obj.(*v1beta1batch.CronJob)
				if !ok {
					return errors.New("could not parse v1beta1batch cronjob definition")
				}
				j.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
				break
			case *v2alpha1batch.CronJob:
				j, ok := obj.(*v2alpha1batch.CronJob)
				if !ok {
					return errors.New("could not parse v2alpha1batch cronjob definition")
				}
				j.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{Name: secretName}}
			}
		}
	}
	return nil
}
