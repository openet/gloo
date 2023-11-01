package contextoptions

import (
	"context"

	"github.com/hashicorp/consul/api"
)

// ContextAccessible is a set of options that are possibly stuffed into the go context
// This is a sub section of cli options that we believe are valuable for use
// in sub cli functions.
type ContextAccessible struct {
	Interactive    bool
	File           string
	Verbose        bool   // currently only used by install and uninstall, sends kubectl command output to terminal
	KubeConfig     string // file to use for kube config, if not standard one.
	ErrorsOnly     bool
	ConfigFilePath string
	Consul         Consul // use consul as config backend
	ReadOnly       bool   // Makes check read only by skipping any checks that create resources in the cluster
	KubeContext    string // kube context to use when interacting with kubernetes
}

type Consul struct {
	UseConsul       bool // enable consul config clients
	RootKey         string
	AllowStaleReads bool
	Client          func() (*api.Client, error)
}

// ContextAccessibleFrom attempts to pull our options that have been stuffed into the go context.
// This relies on "top" being set on the context via the root of the cli package.
func ContextAccessibleFrom(ctx context.Context) ContextAccessible {
	if ctx != nil {
		if contextAccessible, ok := ctx.Value("top").(ContextAccessible); ok {
			return contextAccessible
		}
	}
	return ContextAccessible{}
}

// KubecontextFrom pulls the kube context if it was stuffed into the go context.
// Returns an empty string if it was not set.
func KubecontextFrom(ctx context.Context) string {
	opts := ContextAccessibleFrom(ctx)
	return opts.KubeContext // if context was unset or this value was unset then its empty

}
