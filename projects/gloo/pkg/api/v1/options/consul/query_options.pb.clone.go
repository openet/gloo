// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/query_options.proto

package consul

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *QueryOptions) Clone() proto.Message {
	var target *QueryOptions
	if m == nil {
		return target
	}
	target = &QueryOptions{}

	if h, ok := interface{}(m.GetUseCache()).(clone.Cloner); ok {
		target.UseCache = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.UseCache = proto.Clone(m.GetUseCache()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}
