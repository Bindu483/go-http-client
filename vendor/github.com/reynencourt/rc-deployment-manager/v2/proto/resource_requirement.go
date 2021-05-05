package deploymentManager

import "github.com/reynencourt/rc-common-lib/v2/proto/resource"

func (m *AppResourceRequirements) GetTotalResourceRequirement() *resource.Resources {
	if m.AppResource != nil {
		appRequirement := m.AppResource
		for _, r := range m.ManagedServicesResources {
			if r != nil {
				appRequirement.Add(r)
			}
		}
		return appRequirement
	}
	return nil
}
