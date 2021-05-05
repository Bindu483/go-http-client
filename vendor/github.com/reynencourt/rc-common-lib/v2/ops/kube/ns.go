package kube

import (
	"context"
	"k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Kube) DeleteNamespace(ctx context.Context, namespace string) error {
	return k.clientset.CoreV1().
		Namespaces().
		Delete(ctx, namespace, metav1.DeleteOptions{
			GracePeriodSeconds: &TimeoutSeconds,
		})

}

func (k *Kube) CreateNamespace(ctx context.Context, namespace string) (*v1.Namespace, error) {
	var ns *v1.Namespace
	ns, err := k.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err != nil {
		if apiErrors.IsNotFound(err) {
			nsObj := &v1.Namespace{
				TypeMeta: metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				}}
			ns, err = k.clientset.CoreV1().Namespaces().Create(ctx, nsObj, metav1.CreateOptions{})
			if err != nil {
				return ns, err
			}
		}
	}

	return ns, err
}

func (k *Kube) UpdateNamespace(ctx context.Context, namespace string) (*v1.Namespace, error) {
	ns, err := k.clientset.CoreV1().Namespaces().Update(ctx, &v1.Namespace{
		TypeMeta: metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		}}, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return ns, nil
}

func (k *Kube) ListNamespaces(ctx context.Context) (*v1.NamespaceList, error) {
	namespaces, err := k.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{TimeoutSeconds: &TimeoutSeconds})
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}
