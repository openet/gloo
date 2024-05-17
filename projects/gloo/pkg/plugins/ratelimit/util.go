package ratelimit

import (
	envoycore "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	rlconfig "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v3"
	envoyratelimit "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ratelimit/v3"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func GenerateEnvoyConfigForFilterWith(
	upstreamRef *core.ResourceRef,
	grpcService *ratelimit.GrpcService,
	domain string,
	stage uint32,
	timeout *duration.Duration,
	denyOnFail bool,
	enableXRatelimitHeaders bool,
) *envoyratelimit.RateLimit {

	svc := &envoycore.GrpcService{
		TargetSpecifier: &envoycore.GrpcService_EnvoyGrpc_{
			EnvoyGrpc: &envoycore.GrpcService_EnvoyGrpc{
				ClusterName: translator.UpstreamToClusterName(upstreamRef),
				Authority:   grpcService.GetAuthority(),
			},
		}}

	curtimeout := DefaultTimeout
	if timeout != nil {
		curtimeout = timeout
	}
	xrlHeaders := envoyratelimit.RateLimit_OFF
	if enableXRatelimitHeaders {
		xrlHeaders = envoyratelimit.RateLimit_DRAFT_VERSION_03
	}
	envoyrl := envoyratelimit.RateLimit{
		Domain:                  domain,
		Stage:                   stage,
		RequestType:             RequestType,
		Timeout:                 curtimeout,
		FailureModeDeny:         denyOnFail,
		EnableXRatelimitHeaders: xrlHeaders,

		RateLimitService: &rlconfig.RateLimitServiceConfig{
			TransportApiVersion: envoycore.ApiVersion_V3,
			GrpcService:         svc,
		},
	}
	return &envoyrl
}
