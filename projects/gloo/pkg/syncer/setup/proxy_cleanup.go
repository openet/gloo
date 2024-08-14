package setup

import (
	"context"
	"errors"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	bootstrap "github.com/solo-io/gloo/projects/gloo/pkg/bootstrap/clients"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

func deleteUnusedProxies(ctx context.Context, namespace string, proxyClient v1.ProxyClient) error {
	currentProxies, err := proxyClient.List(namespace, clients.ListOpts{Ctx: ctx})
	if err != nil {
		return err
	}
	deleteErrs := make([]error, 0)
	for _, proxy := range currentProxies {
		if val, ok := proxy.GetMetadata().GetLabels()[utils.ProxyTypeKey]; ok && val == utils.GlooEdgeProxyValue {
			err = proxyClient.Delete(proxy.GetMetadata().GetNamespace(), proxy.GetMetadata().GetName(), clients.DeleteOpts{Ctx: ctx})
			// continue to clean up other proxies
			if err != nil {
				deleteErrs = append(deleteErrs, err)
			}
		}
	}
	if len(deleteErrs) == 0 {
		return nil
	}
	// Concatenate error messages from all the failed deletes
	allErrs := ""
	for _, err := range deleteErrs {
		allErrs += err.Error()
	}
	return errors.New(allErrs)
}
func doProxyCleanup(ctx context.Context, params bootstrap.ConfigFactoryParams, settings *v1.Settings, namespace string) error {
	//Do not clean up proxies if all the resources are held in memory or if proxies are being persisted
	if settings.GetConfigSource() == nil || settings.GetGateway().GetPersistProxySpec().GetValue() {
		return nil
	}
	proxyClientFactory, err := bootstrap.ConfigFactoryForSettings(params, v1.ProxyCrd)
	if err != nil {
		return err
	}
	proxyClient, err := v1.NewProxyClient(ctx, proxyClientFactory)
	if err != nil {
		return err
	}
	if err := proxyClient.Register(); err != nil {
		return err
	}
	return deleteUnusedProxies(ctx, namespace, proxyClient)
}
