package charter

import (
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
	//v1 "k8s.io/kubernetes/staging/src/k8s.io/api/core/v1"
)

//TODO: if an api version that is not supported should appear an error to be thrown
func (c *RcHelmChart) CreateContainerJSON() ([]container.ContainerImageReference, error) {
	var containers []container.ContainerImageReference
	for fileName, k8sObjects := range c.DeploymentManifest.Chart {
		for _, obj := range k8sObjects {
			switch obj.(type) {
			case *corev1.Pod:
				p, ok := obj.(*corev1.Pod)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.InitContainers, &containers)
				break
			case *v1beta1extensions.ReplicaSet:
				p, ok := obj.(*v1beta1extensions.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1apps.ReplicaSet:
				p, ok := obj.(*v1apps.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta2apps.ReplicaSet:
				p, ok := obj.(*v1beta2apps.ReplicaSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *corev1.ReplicationController:
				p, ok := obj.(*corev1.ReplicationController)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1apps.Deployment:
				p, ok := obj.(*v1apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta2apps.Deployment:
				p, ok := obj.(*v1beta2apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta1apps.Deployment:
				p, ok := obj.(*v1beta1apps.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta1extensions.Deployment:
				p, ok := obj.(*v1beta1extensions.Deployment)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1apps.StatefulSet:
				p, ok := obj.(*v1apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta2apps.StatefulSet:
				p, ok := obj.(*v1beta2apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta1apps.StatefulSet:
				p, ok := obj.(*v1beta1apps.StatefulSet)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1apps.DaemonSet:
				return nil, errors.New("daemonsets are not allowed to deploy")
			case *v1beta1extensions.DaemonSet:
				return nil, errors.New("daemonsets are not allowed to deploy")
			case *v1beta2apps.DaemonSet:
				return nil, errors.New("daemonsets are not allowed to deploy")
			case *v1batch.Job:
				p, ok := obj.(*v1batch.Job)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.Template.Spec.InitContainers, &containers)
				break
			case *v1beta1batch.CronJob:
				p, ok := obj.(*v1beta1batch.CronJob)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.JobTemplate.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.JobTemplate.Spec.Template.Spec.InitContainers, &containers)

				break
			case *v2alpha1batch.CronJob:
				p, ok := obj.(*v2alpha1batch.CronJob)
				if !ok {
					return nil, errors.New("could not parse pod definition")
				}
				fetchContainerImageReferences(fileName, p.Spec.JobTemplate.Spec.Template.Spec.Containers, &containers)
				fetchContainerImageReferences(fileName, p.Spec.JobTemplate.Spec.Template.Spec.InitContainers, &containers)
				break
			}
		}
	}

	return containers, nil
}

func fetchContainerImageReferences(filename string, containers []corev1.Container, usedContainers *[]container.ContainerImageReference) {
	for _, c := range containers {
		*usedContainers = append(*usedContainers, container.ContainerImageReference{
			Location:       filename,
			Name:           c.Name,
			ContainerImage: container.ParseImage(c.Image),
		})
	}
}
