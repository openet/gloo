package kubernetes

import (
	"fmt"
	"net/url"

	errors "github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"k8s.io/client-go/kubernetes"
)

var _ discovery.DiscoveryPlugin = new(plugin)

type plugin struct {
	kube kubernetes.Interface

	UpstreamConverter UpstreamConverter

	kubeCoreCache corecache.KubeCoreCache

	settings *v1.Settings
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	kubeSpec, ok := u.UpstreamType.(*v1.Upstream_Kube)
	if !ok {
		return nil, nil
	}

	return url.Parse(fmt.Sprintf("tcp://%v.%v.svc.cluster.local:%v", kubeSpec.Kube.ServiceName, kubeSpec.Kube.ServiceNamespace, kubeSpec.Kube.ServicePort))
}

func NewPlugin(kube kubernetes.Interface, kubeCoreCache corecache.KubeCoreCache) plugins.Plugin {
	return &plugin{
		kube:              kube,
		UpstreamConverter: DefaultUpstreamConverter(),
		kubeCoreCache:     kubeCoreCache,
	}
}

func (p *plugin) Init(params plugins.InitParams) error {
	p.settings = params.Settings
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	// not ours
	kube, ok := in.UpstreamType.(*v1.Upstream_Kube)
	if !ok {
		return nil
	}

	// configure the cluster to use EDS:ADS and call it a day
	xds.SetEdsOnCluster(out, p.settings)

	svcs, err := p.kubeCoreCache.NamespacedServiceLister(kube.Kube.ServiceNamespace).List(labels.NewSelector())
	if err != nil {
		return err
	}
	for _, s := range svcs {
		if s.Name == kube.Kube.ServiceName {
			return nil
		}
	}

	upstreamRef := in.GetMetadata().Ref()
	return errors.Errorf("Upstream %s references the service \"%s\" which does not exist in namespace \"%s\"",
		upstreamRef.String(), kube.Kube.ServiceName, kube.Kube.ServiceNamespace)

}
