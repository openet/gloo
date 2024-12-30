package proxy_syncer

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gateway2/translator/irtranslator"
	ggv2utils "github.com/solo-io/gloo/projects/gateway2/utils"
	"github.com/solo-io/gloo/projects/gateway2/utils/krtutil"
	"github.com/solo-io/go-utils/contextutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/resource"
	"go.uber.org/zap"
	"istio.io/istio/pkg/kube/krt"
)

type uccWithCluster struct {
	Client         ir.UniqlyConnectedClient
	Cluster        envoycache.Resource
	ClusterVersion uint64
	Name           string
	Error          error
}

func (c uccWithCluster) ResourceName() string {
	return fmt.Sprintf("%s/%s", c.Client.ResourceName(), c.Name)
}

func (c uccWithCluster) Equals(in uccWithCluster) bool {
	return c.Client.Equals(in.Client) && c.ClusterVersion == in.ClusterVersion
}

type PerClientEnvoyClusters struct {
	clusters krt.Collection[uccWithCluster]
	index    krt.Index[string, uccWithCluster]
}

func (iu *PerClientEnvoyClusters) FetchClustersForClient(kctx krt.HandlerContext, ucc ir.UniqlyConnectedClient) []uccWithCluster {
	return krt.Fetch(kctx, iu.clusters, krt.FilterIndex(iu.index, ucc.ResourceName()))
}

func NewPerClientEnvoyClusters(
	ctx context.Context,
	krtopts krtutil.KrtOptions,
	translator *irtranslator.UpstreamTranslator,
	upstreams krt.Collection[ir.Upstream],
	uccs krt.Collection[ir.UniqlyConnectedClient],
) PerClientEnvoyClusters {
	ctx = contextutils.WithLogger(ctx, "upstream-translator")
	logger := contextutils.LoggerFrom(ctx).Desugar()

	clusters := krt.NewManyCollection(upstreams, func(kctx krt.HandlerContext, up ir.Upstream) []uccWithCluster {
		logger := logger.With(zap.Stringer("upstream", up))
		uccs := krt.Fetch(kctx, uccs)
		uccWithClusterRet := make([]uccWithCluster, 0, len(uccs))

		for _, ucc := range uccs {
			logger.Debug("applying destination rules for upstream", zap.String("ucc", ucc.ResourceName()))

			c, err := translator.TranslateUpstream(kctx, ucc, up)
			if c == nil {
				continue
			}
			uccWithClusterRet = append(uccWithClusterRet, uccWithCluster{
				Client:         ucc,
				Cluster:        resource.NewEnvoyResource(c),
				Name:           c.GetName(),
				Error:          err,
				ClusterVersion: ggv2utils.HashProto(c),
			})
		}
		return uccWithClusterRet
	}, krtopts.ToOptions("PerClientEnvoyClusters")...)
	idx := krt.NewIndex(clusters, func(ucc uccWithCluster) []string {
		return []string{ucc.Client.ResourceName()}
	})

	return PerClientEnvoyClusters{
		clusters: clusters,
		index:    idx,
	}
}
