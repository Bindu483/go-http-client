package charter

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/container"
	v1apps "k8s.io/api/apps/v1"
	v1beta1apps "k8s.io/api/apps/v1beta1"
	v1beta2apps "k8s.io/api/apps/v1beta2"
	v1batch "k8s.io/api/batch/v1"
	v1beta1batch "k8s.io/api/batch/v1beta1"
	v2alpha1batch "k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	v1beta1extensions "k8s.io/api/extensions/v1beta1"
)

//TODO: if container.json is already updated when we sync, why do we have to this again?
func (c *RcHelmChart) UpdateManifestContainerImageRegistry() error {
	images, err := c.rawChart.GetContainerJSON()
	if err != nil {
		return err
	}
	for _, image := range images {
		registry := image.ContainerImage.Registry
		manifestPath := strings.ReplaceAll(image.Location, fmt.Sprintf("%s/", c.rawChart.Metadata.Name), "")
		for _, obj := range c.DeploymentManifest.Chart[manifestPath] {
			switch obj.(type) {
			case *corev1.Pod:
				p, ok := obj.(*corev1.Pod)
				if !ok {
					return errors.New("could not parse pod definition")
				}
				p.Spec.Containers = updateImageRegistry(registry, &image, p.Spec.Containers)
				p.Spec.InitContainers = updateImageRegistry(registry, &image, p.Spec.InitContainers)
				break
			case *v1beta1extensions.ReplicaSet:
				r, ok := obj.(*v1beta1extensions.ReplicaSet)
				if !ok {
					return errors.New("could not parse replica set definition")
				}
				r.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.Containers)
				r.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.InitContainers)
				break
			case *v1apps.ReplicaSet:
				r, ok := obj.(*v1apps.ReplicaSet)
				if !ok {
					return errors.New("could not parse replica set definition")
				}
				r.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.Containers)
				r.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.InitContainers)
			case *v1beta2apps.ReplicaSet:
				r, ok := obj.(*v1beta2apps.ReplicaSet)
				if !ok {
					return errors.New("could not parse v1beta2apps replica set definition")
				}
				r.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.Containers)
				r.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.InitContainers)
			case *corev1.ReplicationController:
				r, ok := obj.(*corev1.ReplicationController)
				if !ok {
					return errors.New("could not parse corev1 replica set definition")
				}
				r.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.Containers)
				r.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, r.Spec.Template.Spec.InitContainers)
				break
			case *v1apps.Deployment:
				d, ok := obj.(*v1apps.Deployment)
				if !ok {
					return errors.New("could not parse v1apps deployment definition")
				}
				d.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.Containers)
				d.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.InitContainers)
				break
			case *v1beta2apps.Deployment:
				d, ok := obj.(*v1beta2apps.Deployment)
				if !ok {
					return errors.New("could not parse v1beta2apps deployment definition")
				}
				d.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.Containers)
				d.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.InitContainers)
				break
			case *v1beta1apps.Deployment:
				d, ok := obj.(*v1beta1apps.Deployment)
				if !ok {
					return errors.New("could not parse v1beta1apps deployment definition")
				}
				d.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.Containers)
				d.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.InitContainers)
				break
			case *v1beta1extensions.Deployment:
				d, ok := obj.(*v1beta1extensions.Deployment)
				if !ok {
					return errors.New("could not parse v1beta1extensions deployment definition")
				}
				d.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.Containers)
				d.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, d.Spec.Template.Spec.InitContainers)
				break
			case *v1apps.StatefulSet:
				s, ok := obj.(*v1apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1apps statefulset definition")
				}
				s.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.Containers)
				s.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.InitContainers)
				break
			case *v1beta2apps.StatefulSet:
				s, ok := obj.(*v1beta2apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1beta2apps statefulset definition")
				}
				s.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.Containers)
				s.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.InitContainers)
				break
			case *v1beta1apps.StatefulSet:
				s, ok := obj.(*v1beta1apps.StatefulSet)
				if !ok {
					return errors.New("could not parse v1beta1apps statefulset definition")
				}
				s.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.Containers)
				s.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, s.Spec.Template.Spec.InitContainers)
				break
			case *v1apps.DaemonSet:
				return errors.New("daemonsets are ot allowed")
			case *v1beta1extensions.DaemonSet:
				return errors.New("daemonsets are ot allowed")
			case *v1beta2apps.DaemonSet:
				return errors.New("daemonsets are ot allowed")
			case *v1batch.Job:
				j, ok := obj.(*v1batch.Job)
				if !ok {
					return errors.New("could not parse v1batch job definition")
				}
				j.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, j.Spec.Template.Spec.Containers)
				j.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, j.Spec.Template.Spec.InitContainers)
				break
			case *v1beta1batch.CronJob:
				j, ok := obj.(*v1beta1batch.CronJob)
				if !ok {
					return errors.New("could not parse v1beta1batch cronjob definition")
				}
				j.Spec.JobTemplate.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, j.Spec.JobTemplate.Spec.Template.Spec.Containers)
				j.Spec.JobTemplate.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, j.Spec.JobTemplate.Spec.Template.Spec.InitContainers)
				break
			case *v2alpha1batch.CronJob:
				j, ok := obj.(*v2alpha1batch.CronJob)
				if !ok {
					return errors.New("could not parse v2alpha1batch cronjob definition")
				}
				j.Spec.JobTemplate.Spec.Template.Spec.Containers = updateImageRegistry(registry, &image, j.Spec.JobTemplate.Spec.Template.Spec.Containers)
				j.Spec.JobTemplate.Spec.Template.Spec.InitContainers = updateImageRegistry(registry, &image, j.Spec.JobTemplate.Spec.Template.Spec.InitContainers)
			}
		}
	}
	return nil
}

func updateImageRegistry(registry string, image *container.ContainerImageReference, containers []corev1.Container) []corev1.Container {
	updatedContainers := make([]corev1.Container, len(containers))
	for i, c := range containers {
		if c.Name == image.Name {
			image.ContainerImage.Registry = registry
			c.Image = image.ContainerImage.ToString()

		}
		updatedContainers[i] = c
	}
	return updatedContainers
}
