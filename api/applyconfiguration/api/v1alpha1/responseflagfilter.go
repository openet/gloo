// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ResponseFlagFilterApplyConfiguration represents a declarative configuration of the ResponseFlagFilter type for use
// with apply.
type ResponseFlagFilterApplyConfiguration struct {
	Flags []string `json:"flags,omitempty"`
}

// ResponseFlagFilterApplyConfiguration constructs a declarative configuration of the ResponseFlagFilter type for use with
// apply.
func ResponseFlagFilter() *ResponseFlagFilterApplyConfiguration {
	return &ResponseFlagFilterApplyConfiguration{}
}

// WithFlags adds the given value to the Flags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Flags field.
func (b *ResponseFlagFilterApplyConfiguration) WithFlags(values ...string) *ResponseFlagFilterApplyConfiguration {
	for i := range values {
		b.Flags = append(b.Flags, values[i])
	}
	return b
}
