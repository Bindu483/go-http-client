package kube

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	DefaultCpuLimit      = "1000m"
	DefaultMemoryLimit   = "1000M"
	DefaultCpuRequest    = "500m"
	DefaultMemoryRequest = "500M"
)

var LimitRanger = `
apiVersion: v1
kind: LimitRange
metadata:
  name: %s
  namespace: %s
spec:
  limits:
  - default:
      cpu: "%s"
      memory: "%s"
    defaultRequest:
      memory: "%s"
      cpu: "%s"
    type: Container
`

var ErrInvalidRequest = errors.New("Invalid Request")

func (k *Kube) GetDefaultLimitRange(ns string) (*v1.LimitRange, error) {

	var limitRange *v1.LimitRange

	lr := fmt.Sprintf(
		LimitRanger,
		getLimitRangerName(ns),
		ns,
		DefaultCpuLimit,
		DefaultMemoryLimit,
		DefaultMemoryRequest,
		DefaultCpuRequest)

	err := yaml.Unmarshal([]byte(lr), &limitRange)

	if err != nil {
		return nil, err
	}
	return limitRange, nil
}

func getLimitRangerName(ns string) string {
	return fmt.Sprintf("%s-%s", ns, "limitrange")
}

func (k *Kube) CreateLimitRange(ctx context.Context, nameSpace string, limitRange *v1.LimitRange) (*v1.LimitRange, error) {
	if nameSpace == "" {
		return nil, ErrInvalidRequest
	}
	if limitRange == nil {
		return nil, ErrInvalidRequest
	}

	lr, err := k.clientset.CoreV1().LimitRanges(nameSpace).Create(ctx, limitRange, v12.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return lr, nil
}

func (k *Kube) UpdateLimitRange(ctx context.Context, nameSpace string, limitRange *v1.LimitRange) (*v1.LimitRange, error) {
	if nameSpace == "" {
		return nil, ErrInvalidRequest
	}

	if limitRange == nil {
		return nil, ErrInvalidRequest
	}

	lr, err := k.clientset.CoreV1().LimitRanges(nameSpace).Update(ctx, limitRange, v12.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return lr, nil
}

func (k *Kube) DeleteLimitRange(ctx context.Context, ns string) error {
	if ns == "" {
		return ErrInvalidRequest
	}
	return k.clientset.CoreV1().LimitRanges(ns).Delete(ctx, getLimitRangerName(ns), v12.DeleteOptions{})
}
