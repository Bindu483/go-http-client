package deploymentManager

import (
	"github.com/pkg/errors"
)

var (
	ErrRequiredFieldsMissing   = errors.New("some required fields are missing")
	ErrDeployConfigIsNil       = errors.New("deploy app config is nil")
	ErrRequiredValueMissing    = errors.New("required app config missing")
	ErrMSConfigurationMissing  = errors.New("managed services configuration missing")
	ErrMSConfigurationMismatch = errors.New("managed services configuration mismatch")
)

func (d *DeployAppRequest) Validate() error {
	if d.Version == "" || d.SolutionId == "" || d.ClusterId == "" || d.DeploymentId == "" || d.User == "" {
		return ErrRequiredFieldsMissing
	}
	return nil
}
