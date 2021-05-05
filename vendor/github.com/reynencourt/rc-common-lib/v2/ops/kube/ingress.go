package kube

import (
	"context"
	"k8s.io/api/extensions/v1beta1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Kube) ApplyIngress(ctx context.Context, ing *v1beta1.Ingress) error {
	_, err := k.clientset.ExtensionsV1beta1().Ingresses(ing.Namespace).Create(ctx, ing, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *Kube) GetIngress(ctx context.Context, deploymentName string) (*v1beta1.Ingress, error) {
	return k.clientset.ExtensionsV1beta1().Ingresses(deploymentName).Get(ctx, deploymentName, v1.GetOptions{})
}

func (k *Kube) DeleteIngress(ctx context.Context, ns, ingressName string) error {
	return k.clientset.ExtensionsV1beta1().Ingresses(ns).Delete(ctx, ingressName, v1.DeleteOptions{})
}

func (k *Kube) GetIng(ctx context.Context, name, ns string) (*v1beta1.Ingress, error) {
	return k.clientset.ExtensionsV1beta1().Ingresses(ns).Get(ctx, name, v1.GetOptions{})
}

func (k *Kube) CreateIngress(ctx context.Context, ing *v1beta1.Ingress, ns string) error {
	_, err := k.GetIng(ctx, ing.Name, ns)
	if err != nil {
		if apiErrors.IsNotFound(err) {
			_, err = k.clientset.ExtensionsV1beta1().Ingresses(ns).Create(ctx, ing, v1.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}
	return err
}
