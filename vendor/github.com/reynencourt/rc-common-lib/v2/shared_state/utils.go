package shared_state

func dedup(hostEntries []HostEntry) []HostEntry {

	var newHostEntries []HostEntry

	for _, h := range hostEntries {
		for i, newH := range newHostEntries {
			if newH.HostName == h.HostName {
				newHostEntries = append(newHostEntries[:i], newHostEntries[i+1:]...)
			}
		}

		newHostEntries = append(newHostEntries, h)
	}

	return newHostEntries
}

func contains(a string, services []string) bool {

	for _, service := range services {
		if service == a {
			return true
		}
	}

	return false
}

func notApplicableForRemoval(hostEntries []HostEntry, serviceNames []string) []HostEntry {

	var out []HostEntry
	for _, entry := range hostEntries {
		if !contains(entry.HostName, serviceNames) {
			out = append(out, entry)
		}
	}

	return out
}

func notIn(serviceNames []string, hostname string) bool {
	for _, service := range serviceNames {
		if hostname == service {
			return false
		}
	}

	return true
}
