// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FilterTypeApplyConfiguration represents a declarative configuration of the FilterType type for use
// with apply.
type FilterTypeApplyConfiguration struct {
	StatusCodeFilter     *StatusCodeFilterApplyConfiguration   `json:"statusCodeFilter,omitempty"`
	DurationFilter       *DurationFilterApplyConfiguration     `json:"durationFilter,omitempty"`
	NotHealthCheckFilter *bool                                 `json:"notHealthCheckFilter,omitempty"`
	TraceableFilter      *bool                                 `json:"traceableFilter,omitempty"`
	HeaderFilter         *HeaderFilterApplyConfiguration       `json:"headerFilter,omitempty"`
	ResponseFlagFilter   *ResponseFlagFilterApplyConfiguration `json:"responseFlagFilter,omitempty"`
	GrpcStatusFilter     *GrpcStatusFilterApplyConfiguration   `json:"grpcStatusFilter,omitempty"`
	CELFilter            *CELFilterApplyConfiguration          `json:"celFilter,omitempty"`
}

// FilterTypeApplyConfiguration constructs a declarative configuration of the FilterType type for use with
// apply.
func FilterType() *FilterTypeApplyConfiguration {
	return &FilterTypeApplyConfiguration{}
}

// WithStatusCodeFilter sets the StatusCodeFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StatusCodeFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithStatusCodeFilter(value *StatusCodeFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.StatusCodeFilter = value
	return b
}

// WithDurationFilter sets the DurationFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DurationFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithDurationFilter(value *DurationFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.DurationFilter = value
	return b
}

// WithNotHealthCheckFilter sets the NotHealthCheckFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotHealthCheckFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithNotHealthCheckFilter(value bool) *FilterTypeApplyConfiguration {
	b.NotHealthCheckFilter = &value
	return b
}

// WithTraceableFilter sets the TraceableFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TraceableFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithTraceableFilter(value bool) *FilterTypeApplyConfiguration {
	b.TraceableFilter = &value
	return b
}

// WithHeaderFilter sets the HeaderFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeaderFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithHeaderFilter(value *HeaderFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.HeaderFilter = value
	return b
}

// WithResponseFlagFilter sets the ResponseFlagFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResponseFlagFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithResponseFlagFilter(value *ResponseFlagFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.ResponseFlagFilter = value
	return b
}

// WithGrpcStatusFilter sets the GrpcStatusFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GrpcStatusFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithGrpcStatusFilter(value *GrpcStatusFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.GrpcStatusFilter = value
	return b
}

// WithCELFilter sets the CELFilter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CELFilter field is set to the value of the last call.
func (b *FilterTypeApplyConfiguration) WithCELFilter(value *CELFilterApplyConfiguration) *FilterTypeApplyConfiguration {
	b.CELFilter = value
	return b
}
