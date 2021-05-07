package nodeoperations

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/models"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/proto_marshaller"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib"
	"github.com/reynencourt/rc-resource-manager/v2/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

type ListNodeResponse struct {
	Data *ListNodeResponseData `json:"data"`
}

type ListNodeResponseData struct {
	ClusterUnreachable bool                          `json:"cluster_unreachable"`
	Nodes              []*Node                       `json:"nodes"`
	ClusterStatus      resourceManager.ClusterStatus `json:"cluster_status"`
}

type Node struct {
	Capacity         *models.Resources           `json:"capacity"`
	Status           *resourceManager.NodeStatus `json:"status"`
	MemoryPressure   bool                        `json:"memory_pressure"`
	DiskPressure     bool                        `json:"disk_pressure"`
	PidPressure      bool                        `json:"pid_pressure"`
	NodeReady        bool                        `json:"node_ready"`
	NodeName         string                      `json:"node_name"`
	Ip               string                      `json:"ip"`
	NodeType         *resourceManager.NodeType   `json:"node_type"`
	NetworkAvailable bool                        `json:"network_available"`
	TotalAllocatable *models.Resources           `json:"total_allocatable"`
	Allocated        *models.Resources           `json:"allocated"`
}

//Lists nodes for a given cluster
func (nop *NodeOperations) ListNodes() http.Handler {
	return httptransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			listRequest := request.(*resourceManager.ListNodeRequest)

			resp, err := nop.RcManager.ListNodes(ctx, listRequest)
			if err != nil {
				st, _ := status.FromError(err)

				logrus.Error(st.Message())
				return nil, errorlib.Wrap("1", errUnableToListNode)
			}

			return proto_marshaller.Marshal(resp)
		},

		func(_ context.Context, r *http.Request) (interface{}, error) {
			params := mux.Vars(r)
			return &resourceManager.ListNodeRequest{
				ClusterId: params["cluster_id"],
			}, nil
		},
		proto_marshaller.EncodeProtoMessageResponseJSON,
	)
}
