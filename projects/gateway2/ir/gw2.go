package ir

import (
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"google.golang.org/protobuf/types/known/anypb"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// This is the IR that is used in the translation to XDS. it is self contained and no IO/krt is
// needed to process it to xDS.

// As types here are not in krt collections, so no need for equals and resource name.
// Another advantage - because this doesn't appear in any snapshot, we don't need to redact secrets.

type HttpBackend struct {
	Backend Backend
	AttachedPolicies
}

type HttpRouteRuleMatchIR struct {
	ExtensionRefs    AttachedPolicies
	AttachedPolicies AttachedPolicies
	Parent           *HttpRouteIR
	// the rule that delegated to us
	DelegateParent *HttpRouteRuleIR
	HasChildren    bool
	// if there's an error, the gw-api listener to report it in.
	ListenerParentRef gwv1.ParentReference
	Backends          []HttpBackend
	Match             gwv1.HTTPRouteMatch
	MatchIndex        int
	Name              string
}

type ListenerIR struct {
	Name             string
	BindAddress      string
	BindPort         uint32
	AttachedPolicies AttachedPolicies

	HttpFilterChain []HttpFilterChainIR
	TcpFilterChain  []TcpIR
}

type VirtualHost struct {
	Name string
	// technically envoy supports multiple domains per vhost, but gwapi translation doesnt
	// if this changes, we can edit the IR; in the mean time keeping it simple.
	Hostname         string
	Rules            []HttpRouteRuleMatchIR
	AttachedPolicies AttachedPolicies
}

type FilterChainMatch struct {
	SniDomains []string
}
type TlsBundle struct {
	CA            []byte
	PrivateKey    []byte
	CertChain     []byte
	AlpnProtocols []string
}

type FilterChainCommon struct {
	Matcher              FilterChainMatch
	FilterChainName      string
	CustomNetworkFilters []CustomEnvoyFilter
	TLS                  *TlsBundle
}
type CustomEnvoyFilter struct {
	// Determines filter ordering.
	FilterStage plugins.FilterStage[plugins.WellKnownFilterStage]
	// The name of the filter configuration.
	Name string
	// Filter specific configuration.
	Config *anypb.Any
}

type HttpFilterChainIR struct {
	FilterChainCommon
	Vhosts                  []*VirtualHost
	AttachedPolicies        AttachedPolicies
	AttachedNetworkPolicies AttachedPolicies
}

type TcpIR struct {
	FilterChainCommon
	BackendRefs []Backend
}

// this is 1:1 with envoy deployments
// not in a collection so doesn't need a krt interfaces.
type GatewayIR struct {
	Listeners    []ListenerIR
	SourceObject *gwv1.Gateway

	AttachedPolicies     AttachedPolicies
	AttachedHttpPolicies AttachedPolicies
}
