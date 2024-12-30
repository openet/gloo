package setuputils

import (
	"flag"
	"os"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	"k8s.io/apimachinery/pkg/types"
)

const (
	PodNamespace = statusutils.PodNamespaceEnvName
)

var (
	setupNamespace string
	setupName      string
	setupDir       string
)

func SetupNamespaceName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: setupNamespace,
		Name:      setupName,
	}
}

// TODO (ilackarms): move to a flags package
func init() {

	// Allow for more dynamic setting of settings namespace
	// Based on article https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#the-downward-api
	defaultNamespace := os.Getenv(PodNamespace)
	if defaultNamespace == "" {
		defaultNamespace = defaults.GlooSystem
	}
	flag.StringVar(&setupNamespace, "namespace", defaultNamespace, "namespace to watch for settings crd/file")
	flag.StringVar(&setupName, "name", defaults.SettingsName, "name of settings crd/file to use")
	flag.StringVar(&setupDir, "dir", "",
		"directory to find bootstrap settings if not using kubernetes crds")
}
