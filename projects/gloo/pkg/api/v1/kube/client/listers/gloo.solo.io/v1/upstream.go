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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// UpstreamLister helps list Upstreams.
// All objects returned here must be treated as read-only.
type UpstreamLister interface {
	// List lists all Upstreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Upstream, err error)
	// Upstreams returns an object that can list and get Upstreams.
	Upstreams(namespace string) UpstreamNamespaceLister
	UpstreamListerExpansion
}

// upstreamLister implements the UpstreamLister interface.
type upstreamLister struct {
	listers.ResourceIndexer[*v1.Upstream]
}

// NewUpstreamLister returns a new UpstreamLister.
func NewUpstreamLister(indexer cache.Indexer) UpstreamLister {
	return &upstreamLister{listers.New[*v1.Upstream](indexer, v1.Resource("upstream"))}
}

// Upstreams returns an object that can list and get Upstreams.
func (s *upstreamLister) Upstreams(namespace string) UpstreamNamespaceLister {
	return upstreamNamespaceLister{listers.NewNamespaced[*v1.Upstream](s.ResourceIndexer, namespace)}
}

// UpstreamNamespaceLister helps list and get Upstreams.
// All objects returned here must be treated as read-only.
type UpstreamNamespaceLister interface {
	// List lists all Upstreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Upstream, err error)
	// Get retrieves the Upstream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Upstream, error)
	UpstreamNamespaceListerExpansion
}

// upstreamNamespaceLister implements the UpstreamNamespaceLister
// interface.
type upstreamNamespaceLister struct {
	listers.ResourceIndexer[*v1.Upstream]
}
