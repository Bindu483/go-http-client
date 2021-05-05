package deploymentManager

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/rc_spec"
)

var (
	ErrImmutableFieldValueChange = errors.New("immutable field value can not change")
	ErrChangeInManagedServices   = errors.New("change in managed_services section detected")
)

func (req *UpdateDeploymentConfigurationReq) ValidateDeployConfig(spec *rc_spec.RCSpec) (*rc_spec.RCSpec, error) {
	var s rc_spec.RCSpec

	d, err := json.Marshal(spec)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(d, &s)
	if err != nil {
		return nil, err
	}

	if req.Config == nil {
		return nil, ErrDeployConfigIsNil
	}

	if req.Config.AppConfig == nil {
		return nil, ErrDeployConfigIsNil
	}

	appConfig := req.Config.AppConfig.Config

	for _, field := range s.Config.Field {
		val, keyExist := appConfig[field.FieldName]

		if field.Required {
			if !keyExist {
				return nil, ErrRequiredValueMissing
			}
			if field.Immutable {
				if val != field.FieldValue {
					return nil, ErrImmutableFieldValueChange
				}
			}
			//should check for default value against field type
			field.FieldValue = val
		}
		//should check for default value against field type

		if keyExist {
			field.FieldValue = val
		}
	}

	specMSSvcLen := len(s.ManagedServices)
	reqMSSvcLen := len(req.Config.OperatorConfig)

	if specMSSvcLen > 0 {
		if reqMSSvcLen == 0 || specMSSvcLen != reqMSSvcLen {
			return nil, ErrMSConfigurationMissing
		}
	}

	reqMsConfig := req.Config.OperatorConfig
	for ms, config := range s.ManagedServices {
		if len(config.Config) != len(req.Config.OperatorConfig[ms].Config) {
			return nil, ErrMSConfigurationMismatch
		}
		msConfig := reqMsConfig[ms]
		for _, field := range config.Config {
			val, keyExist := msConfig.Config[field.FieldName]

			if field.Required {
				if !keyExist {
					return nil, ErrRequiredValueMissing
				}
				//should check for default value against field type
				field.FieldValue = val
			}
			//should check for default value against field type

			if keyExist {
				field.FieldValue = val
			}
		}
	}

	hasChanged := s.ManagedServicesChanged(&s)
	if hasChanged {
		return nil, ErrChangeInManagedServices
	}

	return &s, nil
}
