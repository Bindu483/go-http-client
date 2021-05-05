package rc_spec

import (
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrUnknownManagedService = errors.New("unsupported managed services")
)

func (m *RCSpec) IsManagedServicesValid() error {
	for k := range m.ManagedServices {
		switch strings.ToUpper(k) {
		case RCManagedService_RCManagedServiceUnknown.String():
		default:
			return ErrUnknownManagedService
		}
	}
	return nil
}

func (m *RCSpec) ManagedServicesChanged(n *RCSpec) bool {
	if len(m.ManagedServices) != len(n.ManagedServices) {
		return true
	}

	for mName, mConfig := range m.ManagedServices {
		nConfig, ok := n.ManagedServices[mName]
		if !ok {
			return true
		}

		if !mConfig.IsSame(nConfig) {
			return true
		}
	}
	return false
}

func (x *OperatorConfig) IsSame(y *OperatorConfig) bool {
	if x.Version != y.Version {
		return false
	}

	if len(x.Config) != len(y.Config) {
		return false
	}

	for i := 0; i < len(x.Config); i++ {
		t := x.Config[i]
		fieldFound := false
		for j := 0; j < len(y.Config); j++ {
			if t.FieldName == y.Config[j].FieldName {
				fieldFound = true
				if t.FieldValue != y.Config[j].FieldValue || t.DefaultValue != y.Config[j].DefaultValue {
					return false
				}
			}
		}
		if !fieldFound {
			return false
		}
	}
	return true
}
