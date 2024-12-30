package translator

import (
	"context"

	"github.com/solo-io/gloo/pkg/utils/syncutil"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap/zapcore"

	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	v1 "github.com/solo-io/gloo/projects/ingress/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

type translatorSyncer struct {
	writeNamespace      string
	writeErrs           chan error
	proxyClient         gloov1.ProxyClient
	ingressClient       v1.IngressClient
	proxyReconciler     gloov1.ProxyReconciler
	requireIngressClass bool

	// support custom ingress class.
	// only relevant when requireIngressClass is true.
	// defaults to 'gloo'
	customIngressClass string

	statusClient resources.StatusClient
}

var (
	// labels used to uniquely identify Proxies that are managed by the Gloo controllers
	proxyLabelsToWrite = map[string]string{
		glooutils.ProxyTypeKey: glooutils.IngressProxyValue,
	}

	// Previously, proxies would be identified with:
	//   created_by: ingress
	// Now, proxies are identified with:
	//   created_by: gloo-ingress
	//
	// We need to ensure that users can successfully upgrade from versions
	// where the previous labels were used, to versions with the new labels.
	// Therefore, we watch Proxies with a superset of the old and new labels, and persist Proxies with new labels.
	//
	// This is only required for backwards compatibility.
	// Once users have upgraded to a version with new labels, we can delete this code and read/write the same labels.
	// gloo-ingress-translator removed in 1.17
	// ingress removed in 1.12
	proxyLabelSelectorOptions = clients.ListOpts{
		ExpressionSelector: glooutils.GetTranslatorSelectorExpression(glooutils.IngressProxyValue, "gloo-ingress-translator", "ingress"),
	}
)

func NewSyncer(writeNamespace string, proxyClient gloov1.ProxyClient, ingressClient v1.IngressClient, writeErrs chan error, requireIngressClass bool, customIngressClass string, statusClient resources.StatusClient) v1.TranslatorSyncer {
	return &translatorSyncer{
		writeNamespace:      writeNamespace,
		writeErrs:           writeErrs,
		proxyClient:         proxyClient,
		ingressClient:       ingressClient,
		proxyReconciler:     gloov1.NewProxyReconciler(proxyClient, statusClient),
		requireIngressClass: requireIngressClass,
		customIngressClass:  customIngressClass,
		statusClient:        statusClient,
	}
}

// TODO (ilackarms): make sure that sync happens if proxies get updated as well; may need to resync
func (s *translatorSyncer) Sync(ctx context.Context, snap *v1.TranslatorSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "translatorSyncer")

	snapHash := hashutils.MustHash(snap)
	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v (%v ingresses)", snapHash,
		len(snap.Ingresses))
	defer logger.Infof("end sync %v", snapHash)

	// stringifying the snapshot may be an expensive operation, so we'd like to avoid building the large
	// string if we're not even going to log it anyway
	if contextutils.GetLogLevel() == zapcore.DebugLevel {
		logger.Debug(syncutil.StringifySnapshot(snap))
	}

	proxy := translateProxy(ctx, s.writeNamespace, snap, s.requireIngressClass, s.customIngressClass)

	var desiredResources gloov1.ProxyList
	if proxy != nil {
		logger.Infof("creating proxy %v", proxy.GetMetadata().Ref())
		proxy.GetMetadata().Labels = proxyLabelsToWrite
		desiredResources = gloov1.ProxyList{proxy}
	}

	proxyTransitionFunction := utils.TransitionFunction(s.statusClient)

	if err := s.proxyReconciler.Reconcile(s.writeNamespace, desiredResources, proxyTransitionFunction, clients.ListOpts{
		Ctx:                ctx,
		Selector:           proxyLabelSelectorOptions.Selector,
		ExpressionSelector: proxyLabelSelectorOptions.ExpressionSelector,
	}); err != nil {
		return err
	}

	return nil
}
