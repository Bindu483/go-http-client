package cb_errors

import "errors"

var (
	ErrNoNodesToRemove       = errors.New("no nodes to removes")
	ErrClusterIdIsMissing    = errors.New("cluster_id is missing")
	ErrDeploymentIdIsMissing = errors.New("deployment_id is missing")
	ErrFailedToDecodeRequest = errors.New("failed to decode request")
	ErrProviderIdIsMissing   = errors.New("provider_id is missing")
	ErrCouldNotDecodeRequest = errors.New("could not decode request")
	ErrSolutionIdIsMissing   = errors.New("solution_id is missing")
	ErrVersionIsMissing      = errors.New("version is missing")
	ErrBackupNameIsMissing   = errors.New("backup_name is missing")
)
