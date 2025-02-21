package namespaces

import (
	"os"
)

// AllNamespaces returns true if the list of namespaces watched is empty or contains only a blank string
// This implies that all namespaces are to be watched
func AllNamespaces(watchNamespaces []string) bool {
	if len(watchNamespaces) == 0 {
		return true
	}
	if len(watchNamespaces) == 1 && watchNamespaces[0] == "" {
		return true
	}
	return false
}

// ProcessWatchNamespaces appends the writeNamespace to the list of watchNamespaces passed if not already present and returns it
func ProcessWatchNamespaces(watchNamespaces []string, writeNamespace string) []string {
	if AllNamespaces(watchNamespaces) {
		return watchNamespaces
	}

	var writeNamespaceProvided bool
	for _, ns := range watchNamespaces {
		if ns == writeNamespace {
			writeNamespaceProvided = true
			break
		}
	}
	if !writeNamespaceProvided {
		watchNamespaces = append(watchNamespaces, writeNamespace)
	}

	return watchNamespaces
}

// GetPodNamespace returns the value of the env var `POD_NAMESPACE` and defaults to `gloo-system` if unset
func GetPodNamespace() string {
	if podNamespace := os.Getenv("POD_NAMESPACE"); podNamespace != "" {
		return podNamespace
	}
	return "gloo-system"
}
