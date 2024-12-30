// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ListenerPolicySpecApplyConfiguration represents a declarative configuration of the ListenerPolicySpec type for use
// with apply.
type ListenerPolicySpecApplyConfiguration struct {
	TargetRef                     *LocalPolicyTargetReferenceApplyConfiguration `json:"targetRef,omitempty"`
	PerConnectionBufferLimitBytes *uint32                                       `json:"perConnectionBufferLimitBytes,omitempty"`
}

// ListenerPolicySpecApplyConfiguration constructs a declarative configuration of the ListenerPolicySpec type for use with
// apply.
func ListenerPolicySpec() *ListenerPolicySpecApplyConfiguration {
	return &ListenerPolicySpecApplyConfiguration{}
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *ListenerPolicySpecApplyConfiguration) WithTargetRef(value *LocalPolicyTargetReferenceApplyConfiguration) *ListenerPolicySpecApplyConfiguration {
	b.TargetRef = value
	return b
}

// WithPerConnectionBufferLimitBytes sets the PerConnectionBufferLimitBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PerConnectionBufferLimitBytes field is set to the value of the last call.
func (b *ListenerPolicySpecApplyConfiguration) WithPerConnectionBufferLimitBytes(value uint32) *ListenerPolicySpecApplyConfiguration {
	b.PerConnectionBufferLimitBytes = &value
	return b
}
