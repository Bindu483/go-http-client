package nodeoperations

import (
	"errors"
)

var (
	errUnableToAddNodes = errors.New("unable to add node. Please contact RC Support")

	errUnableToListNode                     = errors.New("unable to list nodes")
	errUnableToRemoveNodes                  = errors.New("unable to remove nodes")
	errRemoveNodesForceDeleteRequired       = errors.New("unable to remove nodes force delete required")
	errUnableToAddNodesForOngoingOperations = errors.New("unable to add node to cluster due to ongoing operations")
)
