package nodeoperations

import (
	"context"
	"encoding/json"
	"github.com/reynencourt/rc-common-lib/v2/routes"
	"github.com/reynencourt/rc-controller-backend/v2/pkg/configuration"
	"net/http"
)

type NodeOperations struct {
	*configuration.Manager
}

func (nop *NodeOperations) RegisterRoutes() {
	nop.Router.Methods(routes.V2ListNodesRoute.Method).Path(routes.V2ListNodesRoute.Path).Handler(nop.ListNodes())

	nop.Router.Methods(routes.V2AddNodeRoute.Method).Path(routes.V2AddNodeRoute.Path).Handler(nop.AddNodes())

	nop.Router.Methods(routes.V2RemoveNodeRoute.Method).Path(routes.V2RemoveNodeRoute.Path).Handler(nop.RemoveNodes())

	nop.Router.Methods(routes.V2CanRemoveNodeRoute.Method).Path(routes.V2CanRemoveNodeRoute.Path).Handler(nop.CanRemoveNodes())

	nop.Router.Methods(routes.V2GetApprovedInstanceSizesRoute.Method).Path(routes.V2GetApprovedInstanceSizesRoute.Path).Handler(nop.GetApprovedInstanceSizes())
}

func WithRouter(r *configuration.Manager) *NodeOperations {
	return &NodeOperations{r}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
