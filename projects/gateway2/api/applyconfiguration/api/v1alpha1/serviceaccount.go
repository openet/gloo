// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ServiceAccountApplyConfiguration represents a declarative configuration of the ServiceAccount type for use
// with apply.
type ServiceAccountApplyConfiguration struct {
	ExtraLabels      map[string]string `json:"extraLabels,omitempty"`
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
}

// ServiceAccountApplyConfiguration constructs a declarative configuration of the ServiceAccount type for use with
// apply.
func ServiceAccount() *ServiceAccountApplyConfiguration {
	return &ServiceAccountApplyConfiguration{}
}

// WithExtraLabels puts the entries into the ExtraLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ExtraLabels field,
// overwriting an existing map entries in ExtraLabels field with the same key.
func (b *ServiceAccountApplyConfiguration) WithExtraLabels(entries map[string]string) *ServiceAccountApplyConfiguration {
	if b.ExtraLabels == nil && len(entries) > 0 {
		b.ExtraLabels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ExtraLabels[k] = v
	}
	return b
}

// WithExtraAnnotations puts the entries into the ExtraAnnotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ExtraAnnotations field,
// overwriting an existing map entries in ExtraAnnotations field with the same key.
func (b *ServiceAccountApplyConfiguration) WithExtraAnnotations(entries map[string]string) *ServiceAccountApplyConfiguration {
	if b.ExtraAnnotations == nil && len(entries) > 0 {
		b.ExtraAnnotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ExtraAnnotations[k] = v
	}
	return b
}
