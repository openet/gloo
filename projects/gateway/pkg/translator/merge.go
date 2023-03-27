package translator

import (
	"reflect"

	"github.com/imdario/mergo"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Merges the fields of src into dst.
// The fields in dst that have non-zero values will not be overwritten.
func mergeRouteOptions(dst, src *v1.RouteOptions) *v1.RouteOptions {
	if src == nil {
		return dst
	}

	if dst == nil {
		return proto.Clone(src).(*v1.RouteOptions)
	}

	dstValue, srcValue := reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField, srcField := dstValue.Field(i), srcValue.Field(i)
		shallowMerge(dstField, srcField, false)
	}

	return dst
}

// Merges the fields of src into dst.
// The fields in dst that have non-zero values will not be overwritten.
func mergeVirtualHostOptions(dst, src *v1.VirtualHostOptions) *v1.VirtualHostOptions {
	if src == nil {
		return dst
	}

	if dst == nil {
		return proto.Clone(src).(*v1.VirtualHostOptions)
	}

	dstValue, srcValue := reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField, srcField := dstValue.Field(i), srcValue.Field(i)
		shallowMerge(dstField, srcField, false)
	}

	return dst
}

// Sets dst to the value of src, if src is non-zero and dest is zero-valued or overwrite=true.
func shallowMerge(dst, src reflect.Value, overwrite bool) {
	if !src.IsValid() {
		return
	}

	if dst.CanSet() && !isEmptyValue(src) && (overwrite || isEmptyValue(dst)) {
		dst.Set(src)
	}

	return
}

// From src/pkg/encoding/json/encode.go.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isEmptyValue(v.Elem())
	case reflect.Func:
		return v.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

func mergeSslConfig(parent, child *ssl.SslConfig, preventChildOverrides bool) *ssl.SslConfig {
	if child == nil {
		// use parent exactly as-is
		return proto.Clone(parent).(*ssl.SslConfig)
	}
	if parent == nil {
		// use child exactly as-is
		return proto.Clone(child).(*ssl.SslConfig)
	}

	// Clone child to be safe, since we will mutate it
	childClone := proto.Clone(child).(*ssl.SslConfig)
	mergo.Merge(childClone, parent, mergo.WithTransformers(wrapperTransformer{preventChildOverrides}))
	return childClone
}

func mergeHCMSettings(parent, child *hcm.HttpConnectionManagerSettings, preventChildOverrides bool) *hcm.HttpConnectionManagerSettings {
	// Clone to be safe, since we will mutate it
	if child == nil {
		// use parent exactly as-is
		return proto.Clone(parent).(*hcm.HttpConnectionManagerSettings)
	}
	if parent == nil {
		// use child exactly as-is
		return proto.Clone(child).(*hcm.HttpConnectionManagerSettings)
	}

	// Clone child to be safe, since we will mutate it
	childClone := proto.Clone(child).(*hcm.HttpConnectionManagerSettings)
	mergo.Merge(childClone, parent, mergo.WithTransformers(wrapperTransformer{preventChildOverrides}))
	return childClone
}

type wrapperTransformer struct {
	preventChildOverrides bool
}

func (t wrapperTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(wrappers.BoolValue{}) ||
		typ == reflect.TypeOf(wrappers.StringValue{}) ||
		typ == reflect.TypeOf(wrappers.UInt32Value{}) ||
		typ == reflect.TypeOf(duration.Duration{}) ||
		typ == reflect.TypeOf(core.ResourceRef{}) {
		return func(dst, src reflect.Value) error {
			if t.preventChildOverrides {
				dst.Set(src)
			}
			return nil
		}
	}
	return nil
}
