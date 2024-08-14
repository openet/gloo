package consul

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/dynamic_forward_proxy"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func (p *plugin) ProcessRouteAction(
	params plugins.RouteActionParams,
	inAction *v1.RouteAction,
	out *envoy_config_route_v3.RouteAction,
) error {
	switch dest := inAction.GetDestination().(type) {
	case *v1.RouteAction_Single:

		if dest.Single.GetConsul() == nil {
			return nil
		}

		metadataMatch, _, err := getMetadataMatch(dest.Single, params.Snapshot.Upstreams)
		if err != nil {
			return err
		}
		out.MetadataMatch = metadataMatch

		return nil

	case *v1.RouteAction_Multi:
		return setWeightedClusters(params.Params, dest.Multi, out)

	case *v1.RouteAction_UpstreamGroup:
		upstreamGroupRef := dest.UpstreamGroup
		upstreamGroup, err := params.Snapshot.UpstreamGroups.Find(upstreamGroupRef.GetNamespace(), upstreamGroupRef.GetName())
		if err != nil {
			return pluginutils.NewUpstreamGroupNotFoundErr(upstreamGroupRef)
		}
		md := &v1.MultiDestination{
			Destinations: upstreamGroup.GetDestinations(),
		}
		return setWeightedClusters(params.Params, md, out)

	case *v1.RouteAction_ClusterHeader:
		// ClusterHeader must use the naming convention {{clustername}}_{{namespace}}
		out.ClusterSpecifier = &envoy_config_route_v3.RouteAction_ClusterHeader{
			ClusterHeader: inAction.GetClusterHeader(),
		}
		return nil
	case *v1.RouteAction_DynamicForwardProxy:
		out.ClusterSpecifier = &envoy_config_route_v3.RouteAction_Cluster{
			Cluster: dynamic_forward_proxy.GetGeneratedClusterName(params.Listener.GetHttpListener().GetOptions().GetDynamicForwardProxy()),
		}
		return nil
	}
	return eris.Errorf("unknown upstream destination type")
}

func getMetadataMatch(
	dest *v1.Destination,
	allUpstreams v1.UpstreamList,
) (*envoy_config_core_v3.Metadata, *core.ResourceRef, error) {
	usRef, err := upstreams.DestinationToUpstreamRef(dest)
	if err != nil {
		return nil, nil, err
	}

	upstream, err := allUpstreams.Find(usRef.GetNamespace(), usRef.GetName())
	if err != nil {
		return nil, nil, pluginutils.NewUpstreamNotFoundErr(usRef) // should never happen, as we already validated the destination
	}

	return getSubsetMatch(dest, upstream), usRef, nil
}

func setWeightedClusters(params plugins.Params, multiDest *v1.MultiDestination, out *envoy_config_route_v3.RouteAction) error {

	// Index clusters by name so we can look it up by the destination upstream
	clusterMap := make(map[string]*envoy_config_route_v3.WeightedCluster_ClusterWeight)
	for _, weightedCluster := range out.GetWeightedClusters().GetClusters() {
		clusterMap[weightedCluster.GetName()] = weightedCluster
	}

	for _, weightedDest := range multiDest.GetDestinations() {

		if weightedDest.GetDestination().GetConsul() == nil {
			continue
		}

		metadataMatch, usRef, err := getMetadataMatch(weightedDest.GetDestination(), params.Snapshot.Upstreams)
		if err != nil {
			return err
		}

		clusterName := translator.UpstreamToClusterName(usRef)
		correspondentCluster := clusterMap[clusterName]

		correspondentCluster.MetadataMatch = metadataMatch
	}

	return nil
}

func getSubsetMatch(destination *v1.Destination, upstream *v1.Upstream) *envoy_config_core_v3.Metadata {
	var routeMetadata *envoy_config_core_v3.Metadata

	// TODO(marco): consider cleaning up the route API so that subset information is specified on the typed destination
	// If this is a Consul destination, add the correspondent subset information
	// NOTE: if dest.Subset is set on a Consul upstream, this will overwrite it!
	if consulDestination := destination.GetConsul(); consulDestination != nil {
		routeMetadata = consulMetadataMatch(consulDestination, upstream)
	}

	return routeMetadata
}

func consulMetadataMatch(dest *v1.ConsulServiceDestination, upstream *v1.Upstream) *envoy_config_core_v3.Metadata {
	labels := make(map[string]string)

	// If tag filter is provided, set the correspondent metadata.
	// Otherwise don't set them (will match endpoints regardless of tags).
	if len(dest.GetTags()) > 0 {
		labels = BuildTagMetadata(dest.GetTags(), v1.UpstreamList{upstream})
	}

	// If data center filter is provided, set the correspondent metadata.
	// Otherwise don't set them (will match endpoints in any data center).
	if len(dest.GetDataCenters()) > 0 {
		dcLabels := BuildDataCenterMetadata(dest.GetDataCenters(), v1.UpstreamList{upstream})
		for k, v := range dcLabels {
			labels[k] = v
		}
	}

	if len(labels) == 0 {
		return nil
	}

	labelsStruct := &structpb.Struct{
		Fields: map[string]*structpb.Value{},
	}

	for k, v := range labels {
		labelsStruct.GetFields()[k] = &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v,
			},
		}
	}

	return &envoy_config_core_v3.Metadata{
		FilterMetadata: map[string]*structpb.Struct{
			translator.EnvoyLb: labelsStruct,
		},
	}
}
