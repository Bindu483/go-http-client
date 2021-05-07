package nodeoperations

import (
	"context"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/reynencourt/rc-common-lib/v2/cloud_providers/aws"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib"
	"github.com/reynencourt/rc-common-lib/v2/proto/cluster"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/cb_errors"
	"net/http"
	"strconv"
	"strings"
)

type GetApprovedInstanceSizesReq struct {
	Provider cluster.ProviderType `json:"provider"`
}

type GetApprovedInstanceSizesRes struct {
	Provider              cluster.ProviderType `json:"provider"`
	ApprovedInstanceTypes []InstanceSize       `json:"approved_instance_types"`
}

type InstanceSize struct {
	InstanceType string `json:"instance_type"`
	Cpu          int    `json:"cpu"`
	Mem          int    `json:"mem"`
}

func (nop *NodeOperations) GetApprovedInstanceSizes() http.Handler {
	return httptransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*GetApprovedInstanceSizesReq)
			if req.Provider == cluster.ProviderType_ProviderType_Unknown {
				return nil, errorlib.Wrap("1", errors.New("unknown provider"))
			}
			switch req.Provider {
			case cluster.ProviderType_AWS:
				i := aws.GetInstanceTypes()
				var size []InstanceSize
				for _, v := range i {
					size = append(size, InstanceSize{Cpu: v.NoCpu, Mem: v.MemoryCapacity, InstanceType: v.Type})
				}
				return &GetApprovedInstanceSizesRes{Provider: req.Provider, ApprovedInstanceTypes: size}, nil
			}
			return nil, errorlib.Wrap("4", errors.New("unknown provider"))
		},
		func(_ context.Context, r *http.Request) (interface{}, error) {

			var muxVars = mux.Vars(r)

			providerId, ok := muxVars["provider_id"]
			if !ok || strings.TrimSpace(providerId) == "" {
				return nil, cb_errors.ErrProviderIdIsMissing
			}
			providerIdInt, err := strconv.Atoi(providerId)
			if err != nil {
				return nil, errors.New("provider_id should be int")
			}
			return &GetApprovedInstanceSizesReq{Provider: getProvider(providerIdInt)}, nil
		},
		encodeResponse)
}

func getProvider(id int) cluster.ProviderType {
	switch int32(id) {
	case 0:
		return cluster.ProviderType_ProviderType_Unknown
	case 1:
		return cluster.ProviderType_AWS
	case 2:
		return cluster.ProviderType_Azure
	case 3:
		return cluster.ProviderType_OnPrem
	}
	return cluster.ProviderType_ProviderType_Unknown
}
