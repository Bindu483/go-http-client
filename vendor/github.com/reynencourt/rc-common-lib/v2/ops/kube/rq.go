package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Kube) GetResourceQuota(ctx context.Context, name, ns string) (*v1.ResourceQuota, error) {
	return k.clientset.CoreV1().ResourceQuotas(ns).Get(ctx, name, metav1.GetOptions{})
}

func (k *Kube) UpdateResourceQuota(ctx context.Context, r *v1.ResourceQuota) error {
	_, err := k.clientset.CoreV1().ResourceQuotas(r.Namespace).Update(ctx, r, metav1.UpdateOptions{})
	return err
}

func (k *Kube) CreateResourceQuota(ctx context.Context, r *v1.ResourceQuota) error {
	_, err := k.clientset.CoreV1().ResourceQuotas(r.Namespace).Create(ctx, r, metav1.CreateOptions{})
	return err
}
