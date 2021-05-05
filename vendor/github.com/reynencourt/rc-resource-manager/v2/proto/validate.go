package resourceManager

import (
	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/cluster"
)

var (
	ErrProjectIdIsRequired           = errors.New("project_id is a required field")
	ErrUnsupportedProviderType       = errors.New("unsupported provider type")
	ErrClusterNameIsRequired         = errors.New("cluster_name is a required field")
	ErrClusterIdIsRequired           = errors.New("cluster_id is a required field")
	ErrEtcdNodeCountIsZero           = errors.New("state node count is zero")
	ErrWorkerNodeCountIsZero         = errors.New("worker node count is zero")
	ErrMasterNodeCountIsZero         = errors.New("master node count is zero")
	ErrInvalidEtcdNodeCountForHa     = errors.New("state node count should be 1,3,5,... for HA")
	ErrInvalidMasterNodeCountForHa   = errors.New("more than 1 master node is required for HA")
	ErrMinNoOfWorkerNodeCountIsThree = errors.New("min no of worker node count is 3")
	ErrK8sSpecIsNil                  = errors.New("k8s spec is nil")
	ErrClusterIdIsInvalid            = errors.New("cluster id should have only lowercase characters, numbers and hyphens")
	ErrCreatedByIsEmpty              = errors.New("created by field is required")
)

func (x *CreateK8SClusterRequest) Validate() error {
	if x.ProjectId == "" {
		return ErrProjectIdIsRequired
	}
	if x.K8SSpec == nil {
		return ErrK8sSpecIsNil
	}
	return x.K8SSpec.Validate(x.ProviderType)
}

func (x *K8SSpec) Validate(provider cluster.ProviderType) error {
	if x.ClusterName == "" {
		return ErrClusterNameIsRequired
	}
	if x.ClusterId == "" {
		return ErrClusterIdIsRequired
	}
	if err := isValidClusterName(x.ClusterId); err != nil {
		return err
	}
	switch provider {
	case cluster.ProviderType_OnPrem:
		{
			etcdNodeLen := len(x.EtcdInstanceIps)
			masterNodeLen := len(x.MasterInstanceIps)
			workerNodeLen := len(x.WorkerInstanceIps)
			isHa := false
			if etcdNodeLen == 0 {
				return ErrEtcdNodeCountIsZero
			}
			if etcdNodeLen > 1 {
				if etcdNodeLen%2 == 0 {
					return ErrInvalidEtcdNodeCountForHa
				}
				isHa = true
			}
			if workerNodeLen == 0 {
				return ErrWorkerNodeCountIsZero
			}
			if workerNodeLen < 3 {
				return ErrMinNoOfWorkerNodeCountIsThree
			}
			if masterNodeLen == 0 {
				return ErrMasterNodeCountIsZero
			}
			if isHa {
				if masterNodeLen <= 1 {
					return ErrInvalidMasterNodeCountForHa
				}
			}
		}
	case cluster.ProviderType_AWS:
		{
			isHa := false
			if x.EtcdInstanceCount == 0 {
				return ErrEtcdNodeCountIsZero
			}
			if x.EtcdInstanceCount > 1 {
				if x.EtcdInstanceCount%2 == 0 {
					return ErrInvalidEtcdNodeCountForHa
				}
				isHa = true
			}
			if x.WorkerInstanceCount < 3 {
				return ErrMinNoOfWorkerNodeCountIsThree
			}
			if x.MasterInstanceCount == 0 {
				return ErrMasterNodeCountIsZero
			}
			if isHa {
				if x.MasterInstanceCount <= 1 {
					return ErrInvalidMasterNodeCountForHa
				}
			}
		}
	default:
		return ErrUnsupportedProviderType
	}
	return nil
}

func (x *AddNodeToK8SClusterRequest) Validate(provider cluster.ProviderType) error {
	if x.ClusterId == "" {
		return ErrClusterIdIsRequired
	}
	if x.CreatedBy == "" {
		return ErrCreatedByIsEmpty
	}
	return x.K8SSpec.Validate(provider)
}

func (x *AddNodeK8SSpec) Validate(provider cluster.ProviderType) error {
	switch provider {
	case cluster.ProviderType_OnPrem:
		{
			workerNodeLen := len(x.WorkerInstanceIps)
			if workerNodeLen < 0 {
				return ErrWorkerNodeCountIsZero
			}
		}
	case cluster.ProviderType_AWS:
		{
			if x.WorkerInstanceCount < 0 {
				return ErrWorkerNodeCountIsZero
			}
		}
	default:
		return ErrUnsupportedProviderType
	}
	return nil
}

func isValidClusterName(id string) error {
	for i := 0; i < len(id); i++ {
		char := id[i]
		if !(char >= 'a' && char <= 'z' || char >= '0' && char <= '9' || char == '-') {
			return ErrClusterIdIsInvalid
		}
	}
	return nil
}

func (x *MoveClusterReq) Validate() error {
	if x.ClusterId == "" {
		return ErrClusterIdIsRequired
	}
	if x.ProjectId == "" {
		return ErrProjectIdIsRequired
	}
	return nil
}
