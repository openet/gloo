/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSettingses implements SettingsInterface
type FakeSettingses struct {
	Fake *FakeGlooV1
	ns   string
}

var settingsesResource = v1.SchemeGroupVersion.WithResource("settings")

var settingsesKind = v1.SchemeGroupVersion.WithKind("Settings")

// Get takes name of the settings, and returns the corresponding settings object, and an error if there is any.
func (c *FakeSettingses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Settings, err error) {
	emptyResult := &v1.Settings{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(settingsesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Settings), err
}

// List takes label and field selectors, and returns the list of Settingses that match those selectors.
func (c *FakeSettingses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SettingsList, err error) {
	emptyResult := &v1.SettingsList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(settingsesResource, settingsesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SettingsList{ListMeta: obj.(*v1.SettingsList).ListMeta}
	for _, item := range obj.(*v1.SettingsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested settingses.
func (c *FakeSettingses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(settingsesResource, c.ns, opts))

}

// Create takes the representation of a settings and creates it.  Returns the server's representation of the settings, and an error, if there is any.
func (c *FakeSettingses) Create(ctx context.Context, settings *v1.Settings, opts metav1.CreateOptions) (result *v1.Settings, err error) {
	emptyResult := &v1.Settings{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(settingsesResource, c.ns, settings, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Settings), err
}

// Update takes the representation of a settings and updates it. Returns the server's representation of the settings, and an error, if there is any.
func (c *FakeSettingses) Update(ctx context.Context, settings *v1.Settings, opts metav1.UpdateOptions) (result *v1.Settings, err error) {
	emptyResult := &v1.Settings{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(settingsesResource, c.ns, settings, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Settings), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSettingses) UpdateStatus(ctx context.Context, settings *v1.Settings, opts metav1.UpdateOptions) (result *v1.Settings, err error) {
	emptyResult := &v1.Settings{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(settingsesResource, "status", c.ns, settings, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Settings), err
}

// Delete takes name of the settings and deletes it. Returns an error if one occurs.
func (c *FakeSettingses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(settingsesResource, c.ns, name, opts), &v1.Settings{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSettingses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(settingsesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SettingsList{})
	return err
}

// Patch applies the patch and returns the patched settings.
func (c *FakeSettingses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Settings, err error) {
	emptyResult := &v1.Settings{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(settingsesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Settings), err
}
