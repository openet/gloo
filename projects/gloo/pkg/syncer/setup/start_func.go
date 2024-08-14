package setup

import (
	"context"

	"github.com/solo-io/gloo/projects/gloo/pkg/servers/admin"
	"github.com/solo-io/gloo/projects/gloo/pkg/servers/iosnapshot"
	"github.com/solo-io/go-utils/stats"

	"golang.org/x/sync/errgroup"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"

	gateway "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway2/controller"
	"github.com/solo-io/gloo/projects/gateway2/proxy_syncer"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	api "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
)

// StartFunc represents a function that will be called with the initialized bootstrap.Opts
// and Extensions. This is invoked each time the setup_syncer is executed
// (which runs whenever the Setting CR is modified)
type StartFunc func(ctx context.Context, opts bootstrap.Opts, extensions Extensions) error

// ExecuteAsynchronousStartFuncs accepts a collection of StartFunc inputs, and executes them within an Error Group
func ExecuteAsynchronousStartFuncs(
	ctx context.Context,
	opts bootstrap.Opts,
	extensions Extensions,
	startFuncs map[string]StartFunc,
	errorGroup *errgroup.Group,
) {
	for name, start := range startFuncs {
		startFn := start // pike
		namedCtx := contextutils.WithLogger(ctx, name)

		errorGroup.Go(
			func() error {
				contextutils.LoggerFrom(namedCtx).Infof("starting %s goroutine", name)
				err := startFn(namedCtx, opts, extensions)
				if err != nil {
					contextutils.LoggerFrom(namedCtx).Errorf("%s goroutine failed: %v", name, err)
				}
				return err
			},
		)
	}

	contextutils.LoggerFrom(ctx).Debug("main goroutines successfully started")
}

// K8sGatewayControllerStartFunc returns a StartFunc to run the k8s Gateway controller
func K8sGatewayControllerStartFunc(
	proxyClient v1.ProxyClient,
	queueStatusForProxies proxy_syncer.QueueStatusForProxiesFn,
	authConfigClient api.AuthConfigClient,
	routeOptionClient gateway.RouteOptionClient,
	vhOptionClient gateway.VirtualHostOptionClient,
	statusClient resources.StatusClient,
) StartFunc {
	return func(ctx context.Context, opts bootstrap.Opts, extensions Extensions) error {
		statusReporter := reporter.NewReporter(defaults.KubeGatewayReporter, statusClient, routeOptionClient.BaseClient(), vhOptionClient.BaseClient())
		return controller.Start(ctx, controller.StartConfig{
			ExtensionsFactory:         extensions.K8sGatewayExtensionsFactory,
			GlooPluginRegistryFactory: extensions.PluginRegistryFactory,
			Opts:                      opts,
			QueueStatusForProxies:     queueStatusForProxies,

			ProxyClient:             proxyClient,
			AuthConfigClient:        authConfigClient,
			RouteOptionClient:       routeOptionClient,
			VirtualHostOptionClient: vhOptionClient,
			StatusReporter:          statusReporter,

			// Useful for development purposes
			// At the moment, this is not tied to any user-facing API
			Dev: false,
		})
	}
}

// AdminServerStartFunc returns the setup.StartFunc for the Admin Server
// The Admin Server is the groundwork for an Administration Interface, similar to that of Envoy
// https://github.com/solo-io/gloo/issues/6494
// The endpoints that are available on this server are split between two places:
//  1. The default endpoints are defined by our stats server: https://github.com/solo-io/go-utils/blob/8eda16b9878d71673e6a3a9756f6088160f75468/stats/stats.go#L79
//  2. Custom endpoints are defined by our admin server handler in `gloo/pkg/servers/admin`
func AdminServerStartFunc(history iosnapshot.History) StartFunc {
	return func(ctx context.Context, opts bootstrap.Opts, extensions Extensions) error {
		// serverHandlers defines the custom handlers that the Admin Server will support
		serverHandlers := admin.ServerHandlers(ctx, history)

		// The Stats Server is used as the running server for our admin endpoints
		//
		// NOTE: There is a slight difference in how we run this server -vs- how we used to run it
		// In the past, we would start the server once, at the beginning of the running container
		// Now, we start a new server each time we invoke a StartFunc.
		stats.StartCancellableStatsServerWithPort(ctx, stats.DefaultStartupOptions(), serverHandlers)

		return nil
	}
}
