package rmerrors

import "errors"

var (
	ErrClusterIsLocked               = errors.New("cluster is locked")
	ErrClusterCannotBeEmpty          = errors.New("cluster name cannot be empty")
	ErrClusterCreationInProgress     = errors.New("cluster creation in progress or failed")
	ErrClusterDoesNotExist           = errors.New("cluster does not exist")
	ErrClusterWorkspaceDoesNotExist  = errors.New("cluster workspace does not exist")
	ErrStateNotFound                 = errors.New("no state found")
	ErrInitializingWorkspace         = errors.New("could not create workspace")
	ErrClusterIsInvalid              = errors.New("cluster config is invalid")
	ErrCouldNotFetchNodes            = errors.New("could not fetch nodes")
	ErrConnectingToCluster           = errors.New("error connecting to cluster")
	ErrDeleteNodeWillCauseContention = errors.New("deleting node will cause contention")
)
