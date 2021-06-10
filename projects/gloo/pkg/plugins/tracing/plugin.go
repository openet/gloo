package tracing

import (
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoytracing "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/gogo/protobuf/types"
	"github.com/solo-io/gloo/pkg/utils/gogoutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	hcmp "github.com/solo-io/gloo/projects/gloo/pkg/plugins/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/internal/common"
)

// default all tracing percentages to 100%
const oneHundredPercent float32 = 100.0

func NewPlugin() *Plugin {
	return &Plugin{}
}

var _ plugins.Plugin = new(Plugin)
var _ hcmp.HcmPlugin = new(Plugin)
var _ plugins.RoutePlugin = new(Plugin)

type Plugin struct {
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

// Manage the tracing portion of the HCM settings
func (p *Plugin) ProcessHcmSettings(cfg *envoyhttp.HttpConnectionManager, hcmSettings *hcm.HttpConnectionManagerSettings) error {

	// only apply tracing config to the listener is using the HCM plugin
	if hcmSettings == nil {
		return nil
	}

	tracingSettings := hcmSettings.Tracing
	if tracingSettings == nil {
		return nil
	}

	// this plugin will overwrite any prior tracing config
	trCfg := &envoyhttp.HttpConnectionManager_Tracing{}

	customTags := customTags(tracingSettings)
	trCfg.CustomTags = customTags
	trCfg.Verbose = tracingSettings.Verbose

	// Gloo configures envoy as an ingress, rather than an egress
	// 06/2020 removing below- OperationName field is being deprecated, and we set it to the default value anyway
	// trCfg.OperationName = envoyhttp.HttpConnectionManager_Tracing_INGRESS
	if percentages := tracingSettings.GetTracePercentages(); percentages != nil {
		trCfg.ClientSampling = envoySimplePercentWithDefault(percentages.GetClientSamplePercentage(), oneHundredPercent)
		trCfg.RandomSampling = envoySimplePercentWithDefault(percentages.GetRandomSamplePercentage(), oneHundredPercent)
		trCfg.OverallSampling = envoySimplePercentWithDefault(percentages.GetOverallSamplePercentage(), oneHundredPercent)
	} else {
		trCfg.ClientSampling = envoySimplePercent(oneHundredPercent)
		trCfg.RandomSampling = envoySimplePercent(oneHundredPercent)
		trCfg.OverallSampling = envoySimplePercent(oneHundredPercent)
	}
	cfg.Tracing = trCfg
	return nil
}

func customTags(tracingSettings *tracing.ListenerTracingSettings) []*envoytracing.CustomTag {
	var customTags []*envoytracing.CustomTag
	for _, requestHeaderTag := range tracingSettings.RequestHeadersForTags {
		tag := &envoytracing.CustomTag{
			Tag: requestHeaderTag,
			Type: &envoytracing.CustomTag_RequestHeader{
				RequestHeader: &envoytracing.CustomTag_Header{
					Name: requestHeaderTag,
				},
			},
		}
		customTags = append(customTags, tag)
	}
	for _, envVarTag := range tracingSettings.EnvironmentVariablesForTags {
		tag := &envoytracing.CustomTag{
			Tag: envVarTag.Tag,
			Type: &envoytracing.CustomTag_Environment_{
				Environment: &envoytracing.CustomTag_Environment{
					Name:         envVarTag.Name,
					DefaultValue: envVarTag.DefaultValue,
				},
			},
		}
		customTags = append(customTags, tag)
	}
	for _, literalTag := range tracingSettings.LiteralsForTags {
		tag := &envoytracing.CustomTag{
			Tag: literalTag.Tag,
			Type: &envoytracing.CustomTag_Literal_{
				Literal: &envoytracing.CustomTag_Literal{
					Value: literalTag.Value,
				},
			},
		}
		customTags = append(customTags, tag)
	}

	return customTags
}

func envoySimplePercent(numerator float32) *envoy_type.Percent {
	return &envoy_type.Percent{Value: float64(numerator)}
}

// use FloatValue to detect when nil (avoids error-prone float comparisons)
func envoySimplePercentWithDefault(numerator *types.FloatValue, defaultValue float32) *envoy_type.Percent {
	if numerator == nil {
		return envoySimplePercent(defaultValue)
	}
	return envoySimplePercent(numerator.Value)
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoyroute.Route) error {
	if in.Options == nil || in.Options.Tracing == nil {
		return nil
	}
	if percentages := in.GetOptions().GetTracing().TracePercentages; percentages != nil {
		out.Tracing = &envoyroute.Tracing{
			ClientSampling:  common.ToEnvoyPercentageWithDefault(percentages.GetClientSamplePercentage(), oneHundredPercent),
			RandomSampling:  common.ToEnvoyPercentageWithDefault(percentages.GetRandomSamplePercentage(), oneHundredPercent),
			OverallSampling: common.ToEnvoyPercentageWithDefault(percentages.GetOverallSamplePercentage(), oneHundredPercent),
		}
	} else {
		out.Tracing = &envoyroute.Tracing{
			ClientSampling:  common.ToEnvoyv2Percentage(oneHundredPercent),
			RandomSampling:  common.ToEnvoyv2Percentage(oneHundredPercent),
			OverallSampling: common.ToEnvoyv2Percentage(oneHundredPercent),
		}
	}
	descriptor := in.Options.Tracing.RouteDescriptor
	if descriptor != "" {
		out.Decorator = &envoyroute.Decorator{
			Operation: descriptor,
			Propagate: gogoutils.BoolGogoToProto(in.Options.Tracing.GetPropagate()),
		}
	}
	return nil
}
