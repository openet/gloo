package serviceconverter

import (
	"context"
	"strconv"
	"strings"

	"github.com/solo-io/gloo/pkg/utils/envutils"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"google.golang.org/protobuf/types/known/wrapperspb"
	corev1 "k8s.io/api/core/v1"
)

/*
The values for these annotations can be provided in one of two ways:

gloo.solo.io/sslService.secret = mysecret

OR

gloo.solo.io/sslService.secret = 443:mysecret

The former will use ssl on all ports for the service
The latter will use ssl only on port 443 of the service

*/

const GlooSslSecretAnnotation = "gloo.solo.io/sslService.secret"
const GlooSslTlsCertAnnotation = "gloo.solo.io/sslService.tlsCert"
const GlooSslTlsKeyAnnotation = "gloo.solo.io/sslService.tlsKey"
const GlooSslRootCaAnnotation = "gloo.solo.io/sslService.rootCa"
const GlooSslOneWayTlsAnnotation = "gloo.solo.io/sslService.oneWayTls"

// sets UseSsl on the upstream if the service has the relevant port name
type UseSslConverter struct{}

func (u *UseSslConverter) ConvertService(_ context.Context, svc *corev1.Service, port corev1.ServicePort, us *v1.Upstream) error {

	upstreamSslConfig := upstreamSslConfigFromService(svc, port)

	if upstreamSslConfig != nil {
		us.SslConfig = upstreamSslConfig
	}

	return nil
}

func upstreamSslConfigFromService(svc *corev1.Service, svcPort corev1.ServicePort) *ssl.UpstreamSslConfig {

	if svc.Annotations == nil {
		return nil
	}

	// returns empty string if the target port is specified and it's not equal to the serve port
	getAnnotationValue := func(key string) string {
		valWithPort := svc.Annotations[key]

		val, port := splitPortFromValue(valWithPort)
		if port == 0 || port == svcPort.Port {
			return val
		}
		return ""
	}

	secretName := getAnnotationValue(GlooSslSecretAnnotation)
	tlsCert := getAnnotationValue(GlooSslTlsCertAnnotation)
	tlsKey := getAnnotationValue(GlooSslTlsKeyAnnotation)
	rootCa := getAnnotationValue(GlooSslRootCaAnnotation)
	oneWayTls := getAnnotationValue(GlooSslOneWayTlsAnnotation)

	upstreamSslCfg := &ssl.UpstreamSslConfig{}

	if oneWayTls != "" && envutils.IsTruthyValue(oneWayTls) {
		upstreamSslCfg.OneWayTls = &wrapperspb.BoolValue{Value: true}
	} else if oneWayTls != "" {
		upstreamSslCfg.OneWayTls = &wrapperspb.BoolValue{Value: false}
	}

	switch {
	case secretName != "":
		upstreamSslCfg.SslSecrets = &ssl.UpstreamSslConfig_SecretRef{
			SecretRef: &core.ResourceRef{
				Name:      secretName,
				Namespace: svc.Namespace,
			},
		}
	case tlsCert != "" || tlsKey != "" || rootCa != "":
		upstreamSslCfg.SslSecrets = &ssl.UpstreamSslConfig_SslFiles{
			SslFiles: &ssl.SSLFiles{
				TlsCert: tlsCert,
				TlsKey:  tlsKey,
				RootCa:  rootCa,
			},
		}
	default:
		return nil
	}

	return upstreamSslCfg
}

func splitPortFromValue(value string) (string, int32) {
	split := strings.Split(value, ":")
	if len(split) != 2 {
		return value, 0
	}
	i, _ := strconv.Atoi(split[0])
	return split[1], int32(i)
}
