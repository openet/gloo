// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/grpc_service.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_any "github.com/golang/protobuf/ptypes/any"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *GrpcService) Clone() proto.Message {
	var target *GrpcService
	if m == nil {
		return target
	}
	target = &GrpcService{}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if m.GetInitialMetadata() != nil {
		target.InitialMetadata = make([]*HeaderValue, len(m.GetInitialMetadata()))
		for idx, v := range m.GetInitialMetadata() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.InitialMetadata[idx] = h.Clone().(*HeaderValue)
			} else {
				target.InitialMetadata[idx] = proto.Clone(v).(*HeaderValue)
			}

		}
	}

	switch m.TargetSpecifier.(type) {

	case *GrpcService_EnvoyGrpc_:

		if h, ok := interface{}(m.GetEnvoyGrpc()).(clone.Cloner); ok {
			target.TargetSpecifier = &GrpcService_EnvoyGrpc_{
				EnvoyGrpc: h.Clone().(*GrpcService_EnvoyGrpc),
			}
		} else {
			target.TargetSpecifier = &GrpcService_EnvoyGrpc_{
				EnvoyGrpc: proto.Clone(m.GetEnvoyGrpc()).(*GrpcService_EnvoyGrpc),
			}
		}

	case *GrpcService_GoogleGrpc_:

		if h, ok := interface{}(m.GetGoogleGrpc()).(clone.Cloner); ok {
			target.TargetSpecifier = &GrpcService_GoogleGrpc_{
				GoogleGrpc: h.Clone().(*GrpcService_GoogleGrpc),
			}
		} else {
			target.TargetSpecifier = &GrpcService_GoogleGrpc_{
				GoogleGrpc: proto.Clone(m.GetGoogleGrpc()).(*GrpcService_GoogleGrpc),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService_EnvoyGrpc) Clone() proto.Message {
	var target *GrpcService_EnvoyGrpc
	if m == nil {
		return target
	}
	target = &GrpcService_EnvoyGrpc{}

	target.ClusterName = m.GetClusterName()

	target.Authority = m.GetAuthority()

	if h, ok := interface{}(m.GetRetryPolicy()).(clone.Cloner); ok {
		target.RetryPolicy = h.Clone().(*RetryPolicy)
	} else {
		target.RetryPolicy = proto.Clone(m.GetRetryPolicy()).(*RetryPolicy)
	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc{}

	target.TargetUri = m.GetTargetUri()

	if h, ok := interface{}(m.GetChannelCredentials()).(clone.Cloner); ok {
		target.ChannelCredentials = h.Clone().(*GrpcService_GoogleGrpc_ChannelCredentials)
	} else {
		target.ChannelCredentials = proto.Clone(m.GetChannelCredentials()).(*GrpcService_GoogleGrpc_ChannelCredentials)
	}

	if m.GetCallCredentials() != nil {
		target.CallCredentials = make([]*GrpcService_GoogleGrpc_CallCredentials, len(m.GetCallCredentials()))
		for idx, v := range m.GetCallCredentials() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.CallCredentials[idx] = h.Clone().(*GrpcService_GoogleGrpc_CallCredentials)
			} else {
				target.CallCredentials[idx] = proto.Clone(v).(*GrpcService_GoogleGrpc_CallCredentials)
			}

		}
	}

	target.StatPrefix = m.GetStatPrefix()

	target.CredentialsFactoryName = m.GetCredentialsFactoryName()

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*github_com_golang_protobuf_ptypes_struct.Struct)
	}

	if h, ok := interface{}(m.GetPerStreamBufferLimitBytes()).(clone.Cloner); ok {
		target.PerStreamBufferLimitBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PerStreamBufferLimitBytes = proto.Clone(m.GetPerStreamBufferLimitBytes()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetChannelArgs()).(clone.Cloner); ok {
		target.ChannelArgs = h.Clone().(*GrpcService_GoogleGrpc_ChannelArgs)
	} else {
		target.ChannelArgs = proto.Clone(m.GetChannelArgs()).(*GrpcService_GoogleGrpc_ChannelArgs)
	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_SslCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_SslCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_SslCredentials{}

	if h, ok := interface{}(m.GetRootCerts()).(clone.Cloner); ok {
		target.RootCerts = h.Clone().(*DataSource)
	} else {
		target.RootCerts = proto.Clone(m.GetRootCerts()).(*DataSource)
	}

	if h, ok := interface{}(m.GetPrivateKey()).(clone.Cloner); ok {
		target.PrivateKey = h.Clone().(*DataSource)
	} else {
		target.PrivateKey = proto.Clone(m.GetPrivateKey()).(*DataSource)
	}

	if h, ok := interface{}(m.GetCertChain()).(clone.Cloner); ok {
		target.CertChain = h.Clone().(*DataSource)
	} else {
		target.CertChain = proto.Clone(m.GetCertChain()).(*DataSource)
	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_GoogleLocalCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_GoogleLocalCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_GoogleLocalCredentials{}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_ChannelCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_ChannelCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_ChannelCredentials{}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials:

		if h, ok := interface{}(m.GetSslCredentials()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials{
				SslCredentials: h.Clone().(*GrpcService_GoogleGrpc_SslCredentials),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials{
				SslCredentials: proto.Clone(m.GetSslCredentials()).(*GrpcService_GoogleGrpc_SslCredentials),
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault:

		if h, ok := interface{}(m.GetGoogleDefault()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault{
				GoogleDefault: h.Clone().(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault{
				GoogleDefault: proto.Clone(m.GetGoogleDefault()).(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials:

		if h, ok := interface{}(m.GetLocalCredentials()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials{
				LocalCredentials: h.Clone().(*GrpcService_GoogleGrpc_GoogleLocalCredentials),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_ChannelCredentials_LocalCredentials{
				LocalCredentials: proto.Clone(m.GetLocalCredentials()).(*GrpcService_GoogleGrpc_GoogleLocalCredentials),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_CallCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_CallCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_CallCredentials{}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_AccessToken:

		target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_AccessToken{
			AccessToken: m.GetAccessToken(),
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine:

		if h, ok := interface{}(m.GetGoogleComputeEngine()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine{
				GoogleComputeEngine: h.Clone().(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine{
				GoogleComputeEngine: proto.Clone(m.GetGoogleComputeEngine()).(*github_com_golang_protobuf_ptypes_empty.Empty),
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken:

		target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken{
			GoogleRefreshToken: m.GetGoogleRefreshToken(),
		}

	case *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess:

		if h, ok := interface{}(m.GetServiceAccountJwtAccess()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess{
				ServiceAccountJwtAccess: h.Clone().(*GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess{
				ServiceAccountJwtAccess: proto.Clone(m.GetServiceAccountJwtAccess()).(*GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials),
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleIam:

		if h, ok := interface{}(m.GetGoogleIam()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_GoogleIam{
				GoogleIam: h.Clone().(*GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_GoogleIam{
				GoogleIam: proto.Clone(m.GetGoogleIam()).(*GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials),
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_FromPlugin:

		if h, ok := interface{}(m.GetFromPlugin()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_FromPlugin{
				FromPlugin: h.Clone().(*GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_FromPlugin{
				FromPlugin: proto.Clone(m.GetFromPlugin()).(*GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin),
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_StsService_:

		if h, ok := interface{}(m.GetStsService()).(clone.Cloner); ok {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_StsService_{
				StsService: h.Clone().(*GrpcService_GoogleGrpc_CallCredentials_StsService),
			}
		} else {
			target.CredentialSpecifier = &GrpcService_GoogleGrpc_CallCredentials_StsService_{
				StsService: proto.Clone(m.GetStsService()).(*GrpcService_GoogleGrpc_CallCredentials_StsService),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_ChannelArgs) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_ChannelArgs
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_ChannelArgs{}

	if m.GetArgs() != nil {
		target.Args = make(map[string]*GrpcService_GoogleGrpc_ChannelArgs_Value, len(m.GetArgs()))
		for k, v := range m.GetArgs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Args[k] = h.Clone().(*GrpcService_GoogleGrpc_ChannelArgs_Value)
			} else {
				target.Args[k] = proto.Clone(v).(*GrpcService_GoogleGrpc_ChannelArgs_Value)
			}

		}
	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials{}

	target.JsonKey = m.GetJsonKey()

	target.TokenLifetimeSeconds = m.GetTokenLifetimeSeconds()

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials{}

	target.AuthorizationToken = m.GetAuthorizationToken()

	target.AuthoritySelector = m.GetAuthoritySelector()

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin{}

	target.Name = m.GetName()

	switch m.ConfigType.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(clone.Cloner); ok {
			target.ConfigType = &GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig{
				TypedConfig: h.Clone().(*github_com_golang_protobuf_ptypes_any.Any),
			}
		} else {
			target.ConfigType = &GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin_TypedConfig{
				TypedConfig: proto.Clone(m.GetTypedConfig()).(*github_com_golang_protobuf_ptypes_any.Any),
			}
		}

	}

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_CallCredentials_StsService) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_CallCredentials_StsService
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_CallCredentials_StsService{}

	target.TokenExchangeServiceUri = m.GetTokenExchangeServiceUri()

	target.Resource = m.GetResource()

	target.Audience = m.GetAudience()

	target.Scope = m.GetScope()

	target.RequestedTokenType = m.GetRequestedTokenType()

	target.SubjectTokenPath = m.GetSubjectTokenPath()

	target.SubjectTokenType = m.GetSubjectTokenType()

	target.ActorTokenPath = m.GetActorTokenPath()

	target.ActorTokenType = m.GetActorTokenType()

	return target
}

// Clone function
func (m *GrpcService_GoogleGrpc_ChannelArgs_Value) Clone() proto.Message {
	var target *GrpcService_GoogleGrpc_ChannelArgs_Value
	if m == nil {
		return target
	}
	target = &GrpcService_GoogleGrpc_ChannelArgs_Value{}

	switch m.ValueSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_StringValue:

		target.ValueSpecifier = &GrpcService_GoogleGrpc_ChannelArgs_Value_StringValue{
			StringValue: m.GetStringValue(),
		}

	case *GrpcService_GoogleGrpc_ChannelArgs_Value_IntValue:

		target.ValueSpecifier = &GrpcService_GoogleGrpc_ChannelArgs_Value_IntValue{
			IntValue: m.GetIntValue(),
		}

	}

	return target
}
