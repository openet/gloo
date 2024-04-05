// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/connection_limit/connection_limit.proto

package connection_limit

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

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
func (m *ConnectionLimit) Clone() proto.Message {
	var target *ConnectionLimit
	if m == nil {
		return target
	}
	target = &ConnectionLimit{}

	if h, ok := interface{}(m.GetMaxActiveConnections()).(clone.Cloner); ok {
		target.MaxActiveConnections = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxActiveConnections = proto.Clone(m.GetMaxActiveConnections()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetDelayBeforeClose()).(clone.Cloner); ok {
		target.DelayBeforeClose = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.DelayBeforeClose = proto.Clone(m.GetDelayBeforeClose()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	return target
}