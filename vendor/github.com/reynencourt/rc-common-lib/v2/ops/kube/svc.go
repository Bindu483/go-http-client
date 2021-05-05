package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Kube) GetSvc(ctx context.Context, name, ns string) (*v1.Service, error) {
	svc, err := k.clientset.CoreV1().Services(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (k *Kube) CreateSvc(ctx context.Context, svc *v1.Service, ns string) error {
	_, err := k.GetSvc(ctx, svc.Name, ns)
	if err != nil {
		if apiErrors.IsNotFound(err) {
			_, err = k.clientset.CoreV1().Services(ns).Create(ctx, svc, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (k *Kube) GetEp(ctx context.Context, name, ns string) (*v1.Endpoints, error) {
	ep, err := k.clientset.CoreV1().Endpoints(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ep, nil
}

func (k *Kube) CreateEp(ctx context.Context, ep *v1.Endpoints, ns string) error {
	_, err := k.GetEp(ctx, ep.Name, ns)
	if err != nil {
		if apiErrors.IsNotFound(err) {
			_, err = k.clientset.CoreV1().Endpoints(ns).Create(ctx, ep, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}
	return err
}
