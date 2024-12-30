package irtranslator

import (
	"context"
	"fmt"
	"sort"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routerv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	codecv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/upstream_codec/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoytcp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"

	envoy_tls_inspector "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
	"github.com/solo-io/gloo/projects/gateway2/ir"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/go-utils/contextutils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	DefaultHttpStatPrefix  = "http"
	UpstreamCodeFilterName = "envoy.filters.http.upstream_codec"
)

type filterChainTranslator struct {
	listener        ir.ListenerIR
	routeConfigName string

	PluginPass TranslationPassPlugins
}

func computeListenerAddress(bindAddress string, port uint32, reporter reports.GatewayReporter) *envoy_config_core_v3.Address {
	_, isIpv4Address, err := translator.IsIpv4Address(bindAddress)
	if err != nil {
		// TODO: return error ????
		reporter.SetCondition(reports.GatewayCondition{
			Type:    gwv1.GatewayConditionProgrammed,
			Reason:  gwv1.GatewayReasonInvalid,
			Status:  metav1.ConditionFalse,
			Message: "Error processing listener: " + err.Error(),
		})
	}

	return &envoy_config_core_v3.Address{
		Address: &envoy_config_core_v3.Address_SocketAddress{
			SocketAddress: &envoy_config_core_v3.SocketAddress{
				Protocol: envoy_config_core_v3.SocketAddress_TCP,
				Address:  bindAddress,
				PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
					PortValue: port,
				},
				// As of Envoy 1.22: https://www.envoyproxy.io/docs/envoy/latest/version_history/v1.22/v1.22.0.html
				// the Ipv4Compat flag can only be set on Ipv6 address and Ipv4-mapped Ipv6 address.
				// Check if this is a non-padded pure ipv4 address and unset the compat flag if so.
				Ipv4Compat: !isIpv4Address,
			},
		},
	}
}

func tlsInspectorFilter() *envoy_config_listener_v3.ListenerFilter {
	configEnvoy := &envoy_tls_inspector.TlsInspector{}
	msg, _ := anypb.New(configEnvoy)
	return &envoy_config_listener_v3.ListenerFilter{
		Name: wellknown.TlsInspector,
		ConfigType: &envoy_config_listener_v3.ListenerFilter_TypedConfig{
			TypedConfig: msg,
		},
	}
}

func (h *filterChainTranslator) initFilterChain(ctx context.Context, fcc ir.FilterChainCommon, reporter reports.ListenerReporter) *envoy_config_listener_v3.FilterChain {
	info := &FilterChainInfo{
		Match: fcc.Matcher,
		TLS:   fcc.TLS,
	}

	fc := &envoy_config_listener_v3.FilterChain{
		Name:             fcc.FilterChainName,
		FilterChainMatch: info.toMatch(),
		TransportSocket:  info.toTransportSocket(),
	}

	return fc
}
func (h *filterChainTranslator) computeHttpFilters(ctx context.Context, l ir.HttpFilterChainIR, reporter reports.ListenerReporter) []*envoy_config_listener_v3.Filter {
	log := contextutils.LoggerFrom(ctx).Desugar()

	// 1. Generate all the network filters (including the HttpConnectionManager)
	networkFilters, err := h.computeNetworkFiltersForHttp(ctx, l, reporter)
	if err != nil {
		log.DPanic("error computing network filters", zap.Error(err))
		// TODO: report? return error?
		return nil
	}
	if len(networkFilters) == 0 {
		return nil
	}

	return networkFilters
}

func (n *filterChainTranslator) computeNetworkFiltersForHttp(ctx context.Context, l ir.HttpFilterChainIR, reporter reports.ListenerReporter) ([]*envoy_config_listener_v3.Filter, error) {
	hcm := hcmNetworkFilterTranslator{
		routeConfigName: n.routeConfigName,
		PluginPass:      n.PluginPass,
		reporter:        reporter,
	}
	networkFilters := sortNetworkFilters(n.computePreHCMFilters(ctx, l, reporter))
	networkFilter, err := hcm.computeNetworkFilters(ctx, l)
	if err != nil {
		return nil, err
	}
	networkFilters = append(networkFilters, networkFilter)
	return networkFilters, nil
}
func (n *filterChainTranslator) computePreHCMFilters(ctx context.Context, l ir.HttpFilterChainIR, reporter reports.ListenerReporter) []plugins.StagedNetworkFilter {
	var networkFilters []plugins.StagedNetworkFilter
	// Process the network filters.
	for _, plug := range n.PluginPass {
		stagedFilters, err := plug.NetworkFilters(ctx)
		if err != nil {
			reporter.SetCondition(reports.ListenerCondition{
				Type:    gwv1.ListenerConditionProgrammed,
				Reason:  gwv1.ListenerReasonInvalid,
				Status:  metav1.ConditionFalse,
				Message: "Error processing network plugin: " + err.Error(),
			})
			// TODO: return error?
		}

		for _, nf := range stagedFilters {
			if nf.Filter == nil {
				continue
			}
			networkFilters = append(networkFilters, nf)
		}
	}
	networkFilters = append(networkFilters, convertCustomNetworkFilters(l.CustomNetworkFilters)...)
	return networkFilters
}

func convertCustomNetworkFilters(customNetworkFilters []ir.CustomEnvoyFilter) []plugins.StagedNetworkFilter {

	var out []plugins.StagedNetworkFilter
	for _, customFilter := range customNetworkFilters {
		out = append(out, plugins.StagedNetworkFilter{
			Filter: &envoy_config_listener_v3.Filter{
				Name: customFilter.Name,
				ConfigType: &envoy_config_listener_v3.Filter_TypedConfig{
					TypedConfig: customFilter.Config,
				},
			},
			Stage: customFilter.FilterStage,
		})
	}
	return out
}
func sortNetworkFilters(filters plugins.StagedNetworkFilterList) []*envoy_config_listener_v3.Filter {
	sort.Sort(filters)
	var sortedFilters []*envoy_config_listener_v3.Filter
	for _, filter := range filters {
		sortedFilters = append(sortedFilters, filter.Filter)
	}
	return sortedFilters
}

type hcmNetworkFilterTranslator struct {
	routeConfigName string
	PluginPass      TranslationPassPlugins
	reporter        reports.ListenerReporter
	listener        ir.HttpFilterChainIR
}

func (h *hcmNetworkFilterTranslator) computeNetworkFilters(ctx context.Context, l ir.HttpFilterChainIR) (*envoy_config_listener_v3.Filter, error) {
	ctx = contextutils.WithLogger(ctx, "compute_http_connection_manager")

	// 1. Initialize the HCM
	httpConnectionManager := h.initializeHCM()

	// 2. Apply HttpFilters
	var err error
	httpConnectionManager.HttpFilters = h.computeHttpFilters(ctx, l)

	// 3. Allow any HCM plugins to make their changes, with respect to any changes the core plugin made
	//	for _, hcmPlugin := range h.hcmPlugins {
	//		if err := hcmPlugin.ProcessHcmNetworkFilter(params, h.parentListener, h.listener, httpConnectionManager); err != nil {
	//			h.reporter.SetCondition(reports.ListenerCondition{
	//				Type:    gwv1.ListenerConditionProgrammed,
	//				Reason:  gwv1.ListenerReasonInvalid,
	//				Status:  metav1.ConditionFalse,
	//				Message: "Error processing HCM plugin: " + err.Error(),
	//			})
	//		}
	//	}
	// TODO: should we enable websockets by default?

	// 4. Generate the typedConfig for the HCM
	hcmFilter, err := NewFilterWithTypedConfig(wellknown.HTTPConnectionManager, httpConnectionManager)
	if err != nil {
		contextutils.LoggerFrom(ctx).DPanic("failed to convert proto message to struct")
		return nil, fmt.Errorf("failed to convert proto message to any: %w", err)
	}

	return hcmFilter, nil
}

func (h *hcmNetworkFilterTranslator) initializeHCM() *envoyhttp.HttpConnectionManager {
	statPrefix := h.listener.FilterChainName
	if statPrefix == "" {
		statPrefix = DefaultHttpStatPrefix
	}

	return &envoyhttp.HttpConnectionManager{
		CodecType:     envoyhttp.HttpConnectionManager_AUTO,
		StatPrefix:    statPrefix,
		NormalizePath: wrapperspb.Bool(true),
		MergeSlashes:  true,
		RouteSpecifier: &envoyhttp.HttpConnectionManager_Rds{
			Rds: &envoyhttp.Rds{
				ConfigSource: &envoy_config_core_v3.ConfigSource{
					ResourceApiVersion: envoy_config_core_v3.ApiVersion_V3,
					ConfigSourceSpecifier: &envoy_config_core_v3.ConfigSource_Ads{
						Ads: &envoy_config_core_v3.AggregatedConfigSource{},
					},
				},
				RouteConfigName: h.routeConfigName,
			},
		},
	}
}

func (h *hcmNetworkFilterTranslator) computeHttpFilters(ctx context.Context, l ir.HttpFilterChainIR) []*envoyhttp.HttpFilter {
	var httpFilters plugins.StagedHttpFilterList

	log := contextutils.LoggerFrom(ctx).Desugar()

	// run the HttpFilter Plugins
	for _, plug := range h.PluginPass {
		stagedFilters, err := plug.HttpFilters(ctx, l.FilterChainCommon)
		if err != nil {
			// what to do with errors here? ignore the listener??
			h.reporter.SetCondition(reports.ListenerCondition{
				Type:    gwv1.ListenerConditionProgrammed,
				Reason:  gwv1.ListenerReasonInvalid,
				Status:  metav1.ConditionFalse,
				Message: "Error processing http plugin: " + err.Error(),
			})
			// TODO: return false?
		}

		for _, httpFilter := range stagedFilters {
			if httpFilter.Filter == nil {
				log.Warn("HttpFilters() returned nil", zap.String("name", plug.Name))
				continue
			}
			httpFilters = append(httpFilters, httpFilter)
		}
	}
	//	httpFilters = append(httpFilters, CustomHttpFilters(h.listener)...)

	// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_filters#filter-ordering
	// HttpFilter ordering determines the order in which the HCM will execute the filter.

	// 1. Sort filters by stage
	// "Stage" is the type we use to specify when a filter should be run
	envoyHttpFilters := sortHttpFilters(httpFilters)

	// 2. Configure the router filter
	// As outlined by the Envoy docs, the last configured filter has to be a terminal filter.
	// We set the Router filter (https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#config-http-filters-router)
	// as the terminal filter in k8sgateway.
	routerV3 := routerv3.Router{}

	h.computeUpstreamHTTPFilters(ctx, l, &routerV3)

	//	// TODO it would be ideal of SuppressEnvoyHeaders and DynamicStats could be moved out of here set
	//	// in a separate router plugin
	//	if h.listener.GetOptions().GetRouter().GetSuppressEnvoyHeaders().GetValue() {
	//		routerV3.SuppressEnvoyHeaders = true
	//	}
	//
	//	routerV3.DynamicStats = h.listener.GetOptions().GetRouter().GetDynamicStats()

	newStagedFilter, err := plugins.NewStagedFilter(
		wellknown.Router,
		&routerV3,
		plugins.AfterStage(plugins.RouteStage),
	)
	if err != nil {
		h.reporter.SetCondition(reports.ListenerCondition{
			Type:    gwv1.ListenerConditionProgrammed,
			Reason:  gwv1.ListenerReasonInvalid,
			Status:  metav1.ConditionFalse,
			Message: "Error processing http plugins: " + err.Error(),
		})
		// TODO: return false?
	}

	envoyHttpFilters = append(envoyHttpFilters, newStagedFilter.Filter)

	return envoyHttpFilters
}

func (h *hcmNetworkFilterTranslator) computeUpstreamHTTPFilters(ctx context.Context, l ir.HttpFilterChainIR, routerV3 *routerv3.Router) {
	upstreamHttpFilters := plugins.StagedUpstreamHttpFilterList{}
	log := contextutils.LoggerFrom(ctx).Desugar()
	for _, plug := range h.PluginPass {
		stagedFilters, err := plug.UpstreamHttpFilters(ctx)
		if err != nil {
			// what to do with errors here? ignore the listener??
			h.reporter.SetCondition(reports.ListenerCondition{
				Type:    gwv1.ListenerConditionProgrammed,
				Reason:  gwv1.ListenerReasonInvalid,
				Status:  metav1.ConditionFalse,
				Message: "Error processing upstream http plugin: " + err.Error(),
			})
			// TODO: return false?
		}
		for _, httpFilter := range stagedFilters {
			if httpFilter.Filter == nil {
				log.Warn("HttpFilters() returned nil", zap.String("name", plug.Name))
				continue
			}
			upstreamHttpFilters = append(upstreamHttpFilters, httpFilter)
		}
	}

	if len(upstreamHttpFilters) == 0 {
		return
	}

	sort.Sort(upstreamHttpFilters)

	sortedFilters := make([]*envoyhttp.HttpFilter, len(upstreamHttpFilters))
	for i, filter := range upstreamHttpFilters {
		sortedFilters[i] = filter.Filter
	}

	msg, err := anypb.New(&codecv3.UpstreamCodec{})
	if err != nil {
		// what to do with errors here? ignore the listener??
		h.reporter.SetCondition(reports.ListenerCondition{
			Type:    gwv1.ListenerConditionProgrammed,
			Reason:  gwv1.ListenerReasonInvalid,
			Status:  metav1.ConditionFalse,
			Message: "failed to convert proto message to any: " + err.Error(),
		})
		return
	}

	routerV3.UpstreamHttpFilters = sortedFilters
	routerV3.UpstreamHttpFilters = append(routerV3.GetUpstreamHttpFilters(), &envoyhttp.HttpFilter{
		Name: UpstreamCodeFilterName,
		ConfigType: &envoyhttp.HttpFilter_TypedConfig{
			TypedConfig: msg,
		},
	})
}

func sortHttpFilters(filters plugins.StagedHttpFilterList) []*envoyhttp.HttpFilter {
	sort.Sort(filters)
	var sortedFilters []*envoyhttp.HttpFilter
	for _, filter := range filters {
		sortedFilters = append(sortedFilters, filter.Filter)
	}
	return sortedFilters
}

func (h *filterChainTranslator) computeTcpFilters(ctx context.Context, l ir.TcpIR, reporter reports.ListenerReporter) []*envoy_config_listener_v3.Filter {
	networkFilters := sortNetworkFilters(h.computeNetworkFiltersForTcp(l))

	cfg := &envoytcp.TcpProxy{
		StatPrefix: l.FilterChainName,
	}
	if len(l.BackendRefs) == 1 {
		cfg.ClusterSpecifier = &envoytcp.TcpProxy_Cluster{
			Cluster: l.BackendRefs[0].ClusterName,
		}
	} else {
		var wc envoytcp.TcpProxy_WeightedCluster
		for _, route := range l.BackendRefs {
			w := route.Weight
			if w == 0 {
				w = 1
			}
			wc.Clusters = append(wc.GetClusters(), &envoytcp.TcpProxy_WeightedCluster_ClusterWeight{
				Name:   route.ClusterName,
				Weight: w,
			})
		}
		cfg.ClusterSpecifier = &envoytcp.TcpProxy_WeightedClusters{
			WeightedClusters: &wc,
		}
	}

	tcpFilter, _ := NewFilterWithTypedConfig(wellknown.TCPProxy, cfg)

	return append(networkFilters, tcpFilter)
}

func (t *filterChainTranslator) computeNetworkFiltersForTcp(l ir.TcpIR) []plugins.StagedNetworkFilter {
	var networkFilters []plugins.StagedNetworkFilter
	// Process the network filters.
	//for _, plug := range t.networkPlugins {
	//	stagedFilters, err := plug.NetworkFiltersTCP(params, t.listener)
	//	if err != nil {
	//		validation.AppendTCPListenerError(t.report, validationapi.TcpListenerReport_Error_ProcessingError, err.Error())
	//	}
	//
	//	for _, nf := range stagedFilters {
	//		if nf.Filter == nil {
	//			log.Warnf("plugin %v implements NetworkFilters() but returned nil", plug.Name())
	//			continue
	//		}
	//		networkFilters = append(networkFilters, nf)
	//	}
	//}

	networkFilters = append(networkFilters, convertCustomNetworkFilters(l.CustomNetworkFilters)...)
	return networkFilters
}

func NewFilterWithTypedConfig(name string, config proto.Message) (*envoy_config_listener_v3.Filter, error) {
	s := &envoy_config_listener_v3.Filter{
		Name: name,
	}

	if config != nil {
		marshalledConf, err := anypb.New(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return &envoy_config_listener_v3.Filter{}, err
		}

		s.ConfigType = &envoy_config_listener_v3.Filter_TypedConfig{
			TypedConfig: marshalledConf,
		}
	}

	return s, nil
}

type SslConfig struct {
	Bundle     TlsBundle
	SniDomains []string
}
type TlsBundle struct {
	CA         []byte
	PrivateKey []byte
	CertChain  []byte
}

type FilterChainInfo struct {
	Match ir.FilterChainMatch
	TLS   *ir.TlsBundle
}

func (info *FilterChainInfo) toMatch() *envoy_config_listener_v3.FilterChainMatch {
	if info == nil {
		return nil
	}

	// right now only sni domains is in the match, so if empty, return a nil match
	if len(info.Match.SniDomains) == 0 {
		return nil
	}

	return &envoy_config_listener_v3.FilterChainMatch{
		ServerNames: info.Match.SniDomains,
	}
}

func (info *FilterChainInfo) toTransportSocket() *envoy_config_core_v3.TransportSocket {
	if info == nil {
		return nil
	}
	ssl := info.TLS
	if ssl == nil {
		return nil
	}

	common := &envoyauth.CommonTlsContext{
		// default params
		TlsParams:     &envoyauth.TlsParameters{},
		AlpnProtocols: ssl.AlpnProtocols,
	}

	common.TlsCertificates = []*envoyauth.TlsCertificate{
		{
			CertificateChain: bytesDataSource(ssl.CertChain),
			PrivateKey:       bytesDataSource(ssl.PrivateKey),
		},
	}

	//	var requireClientCert *wrappers.BoolValue
	//	if common.GetValidationContextType() != nil {
	//		requireClientCert = &wrappers.BoolValue{Value: !dc.GetOneWayTls().GetValue()}
	//	}

	// default alpn for downstreams.
	//	if len(common.GetAlpnProtocols()) == 0 {
	//		common.AlpnProtocols = []string{"h2", "http/1.1"}
	//	} else if len(common.GetAlpnProtocols()) == 1 && common.GetAlpnProtocols()[0] == AllowEmpty { // allow override for advanced usage to set to a dangerous setting
	//		common.AlpnProtocols = []string{}
	//	}

	out := &envoyauth.DownstreamTlsContext{
		CommonTlsContext: common,
	}
	typedConfig, _ := anypb.New(out)

	return &envoy_config_core_v3.TransportSocket{
		Name:       wellknown.TransportSocketTls,
		ConfigType: &envoy_config_core_v3.TransportSocket_TypedConfig{TypedConfig: typedConfig},
	}
}
func bytesDataSource(s []byte) *envoy_config_core_v3.DataSource {
	return &envoy_config_core_v3.DataSource{
		Specifier: &envoy_config_core_v3.DataSource_InlineBytes{
			InlineBytes: s,
		},
	}

}
