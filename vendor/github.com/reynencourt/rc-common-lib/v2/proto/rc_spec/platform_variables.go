package rc_spec

func (m *RCSpec) ReplacePlatformVariables(platformVars map[string]string) {
	for k, v := range m.PlatformVariables {
		if value, match := platformVars[v]; match {
			m.PlatformVariables[k] = value
		}
	}
}
