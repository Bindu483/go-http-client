package nodeoperations

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib/rmerrors"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/cb_errors"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/proto_marshaller"
	resourceManager "github.com/reynencourt/rc-resource-manager/v2/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type CanRemoveNodeReq struct {
	ClusterId string   `json:"cluster_id"`
	NodeNames []string `json:"node_names"`
}

func (req *CanRemoveNodeReq) Validate() error {
	if len(req.NodeNames) == 0 {
		return cb_errors.ErrNoNodesToRemove
	}
	if strings.TrimSpace(req.ClusterId) == "" {
		return cb_errors.ErrClusterIdIsMissing
	}
	return nil
}

func (nop *NodeOperations) CanRemoveNodes() http.Handler {
	return httptransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {

			req, ok := request.(*CanRemoveNodeReq)
			if !ok {
				return nil, cb_errors.ErrFailedToDecodeRequest
			}

			resp, err := nop.RcManager.CanDeleteNode(context.Background(), &resourceManager.CanDeleteNodeReq{
				NodeNames: req.NodeNames,
				ClusterId: req.ClusterId,
			})

			if err != nil {
				logrus.WithError(err).Error("error happening when calling CanDeleteNode")
				st, ok := status.FromError(err)
				if ok {
					if st.Code() == codes.FailedPrecondition {
						return nil, errorlib.Wrap("CB101", errRemoveNodesForceDeleteRequired)
					} else if st.Message() == rmerrors.ErrClusterIsLocked.Error() {
						return nil, errorlib.Wrap("CLOP0006", errUnableToAddNodesForOngoingOperations)
					} else {
						logrus.Error(st.Message())
						return nil, errorlib.Wrap("CB102", errUnableToRemoveNodes)
					}
				} else {
					logrus.Error(st.Message())
					return nil, errorlib.Wrap("CB102", errUnableToRemoveNodes)
				}
			}
			return proto_marshaller.Marshal(resp)
		},
		func(_ context.Context, r *http.Request) (interface{}, error) {
			params := mux.Vars(r)
			var canRemoveNodeReq CanRemoveNodeReq
			err := json.NewDecoder(r.Body).Decode(&canRemoveNodeReq)
			if err != nil {
				logrus.WithError(err).Error("failed to decode can remove node request")
				return nil, err
			}
			clusterId := params["cluster_id"]
			if clusterId == "" {
				return nil, errors.New("cluster id is a required field")
			}

			canRemoveNodeReq.ClusterId = clusterId

			err = canRemoveNodeReq.Validate()
			if err != nil {
				logrus.WithError(err).Error("request validation failed")
				return nil, err
			}

			return &canRemoveNodeReq, err
		},
		proto_marshaller.EncodeProtoMessageResponseJSON,
	)
}
