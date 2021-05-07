package nodeoperations

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib"
	"github.com/reynencourt/rc-common-lib/v2/ops/errorlib/rmerrors"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/cb_errors"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/router/route/proto_marshaller"
	resourceManager "github.com/reynencourt/rc-resource-manager/v2/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type RemoveNodeReq struct {
	NodeNames []string `json:"node_names"`
	Force     bool     `json:"force"`
}

func (nop *NodeOperations) RemoveNodes() http.Handler {
	return httptransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {

			req := request.(*resourceManager.RemoveNodeRequest)

			resp, err := nop.RcManager.RemoveNodes(ctx, req)
			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					if st.Code() == codes.FailedPrecondition {
						return nil, errorlib.Wrap("CB101", errRemoveNodesForceDeleteRequired)
					} else if st.Message() == rmerrors.ErrClusterIsLocked.Error() {
						return nil, errorlib.Wrap("CLOP0006", errUnableToAddNodesForOngoingOperations)
					} else if st.Message() == rmerrors.ErrDeleteNodeWillCauseContention.Error() {
						return nil, errorlib.Wrap("CLOP0006", rmerrors.ErrDeleteNodeWillCauseContention)
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
			var removeNodeReq RemoveNodeReq
			err := json.NewDecoder(r.Body).Decode(&removeNodeReq)
			if err != nil {
				logrus.WithError(err).Error("failed to decode add node request")
				return nil, err
			}
			clusterId := params["cluster_id"]
			if clusterId == "" {
				return nil, errors.New("cluster id is a required field")
			}

			err = removeNodeReq.Validate()
			if err != nil {
				logrus.WithError(err).Error("request validation failed")
				return nil, err
			}

			return &resourceManager.RemoveNodeRequest{
				ClusterId:   clusterId,
				NodeNames:   removeNodeReq.NodeNames,
				ForceRemove: removeNodeReq.Force,
			}, err
		},
		proto_marshaller.EncodeProtoMessageResponseJSON,
	)
}

func (req *RemoveNodeReq) Validate() error {
	if len(req.NodeNames) == 0 {
		return cb_errors.ErrNoNodesToRemove
	}
	return nil
}
