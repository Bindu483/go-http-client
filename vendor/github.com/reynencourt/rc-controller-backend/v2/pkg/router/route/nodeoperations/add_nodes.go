package nodeoperations

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib/rmerrors"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/proto_marshaller"
	"github.com/reynencourt/rc-resource-manager/v2/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"net/http"
)

type AddNodeRequest struct {
	CreatedBy           string   `json:"created_by"`
	WorkerIps           []string `json:"worker_ips"`
	WorkerInstanceCount uint32   `json:"worker_instance_count"`
	WorkerInstanceSize  string   `json:"worker_instance_size"`
	StorageSize         uint32   `json:"storage_size"`
}

//AddNodes rest call for adding nodes to the particular cluster
func (nop *NodeOperations) AddNodes() http.Handler {
	return httptransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {

			req := request.(*resourceManager.AddNodeToK8SClusterRequest)

			response, err := nop.RcManager.AddNodes(ctx, req)
			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					logrus.WithField("add_node", "add_node").Error(st.Message())
					if st.Message() == rmerrors.ErrClusterIsLocked.Error() {
						return nil, errorlib.Wrap("CLOP0006", errUnableToAddNodesForOngoingOperations)
					}
				}
				logrus.WithError(err).Error("failed to add nodes")
				return nil, errorlib.Wrap("1", errUnableToAddNodes)
			}
			return proto_marshaller.Marshal(response)
		},
		func(_ context.Context, r *http.Request) (interface{}, error) {
			params := mux.Vars(r)

			var addNodeReq AddNodeRequest
			err := json.NewDecoder(r.Body).Decode(&addNodeReq)
			if err != nil {
				logrus.WithError(err).Error("failed to decode add node request")
				return nil, err
			}

			clusterId := params["cluster_id"]
			if clusterId == "" {
				return nil, errors.New("cluster id is a required field")
			}

			return &resourceManager.AddNodeToK8SClusterRequest{
				ClusterId: clusterId,
				K8SSpec: &resourceManager.AddNodeK8SSpec{
					WorkerInstanceCount: addNodeReq.WorkerInstanceCount,
					InstanceSize:        addNodeReq.WorkerInstanceSize,
					StorageSize:         addNodeReq.StorageSize,
					WorkerInstanceIps:   addNodeReq.WorkerIps,
				},
				CreatedBy: addNodeReq.CreatedBy,
			}, err
		},
		proto_marshaller.EncodeProtoMessageResponseJSON,
	)
}
