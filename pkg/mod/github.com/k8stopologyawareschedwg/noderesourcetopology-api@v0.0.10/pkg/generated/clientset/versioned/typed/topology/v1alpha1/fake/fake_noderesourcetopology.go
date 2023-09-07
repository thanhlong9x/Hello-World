/*
Copyright 2020 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeResourceTopologies implements NodeResourceTopologyInterface
type FakeNodeResourceTopologies struct {
	Fake *FakeTopologyV1alpha1
	ns   string
}

var noderesourcetopologiesResource = schema.GroupVersionResource{Group: "topology.node.k8s.io", Version: "v1alpha1", Resource: "noderesourcetopologies"}

var noderesourcetopologiesKind = schema.GroupVersionKind{Group: "topology.node.k8s.io", Version: "v1alpha1", Kind: "NodeResourceTopology"}

// Get takes name of the nodeResourceTopology, and returns the corresponding nodeResourceTopology object, and an error if there is any.
func (c *FakeNodeResourceTopologies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeResourceTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(noderesourcetopologiesResource, c.ns, name), &v1alpha1.NodeResourceTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeResourceTopology), err
}

// List takes label and field selectors, and returns the list of NodeResourceTopologies that match those selectors.
func (c *FakeNodeResourceTopologies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeResourceTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(noderesourcetopologiesResource, noderesourcetopologiesKind, c.ns, opts), &v1alpha1.NodeResourceTopologyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeResourceTopologyList{ListMeta: obj.(*v1alpha1.NodeResourceTopologyList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeResourceTopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeResourceTopologies.
func (c *FakeNodeResourceTopologies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(noderesourcetopologiesResource, c.ns, opts))

}

// Create takes the representation of a nodeResourceTopology and creates it.  Returns the server's representation of the nodeResourceTopology, and an error, if there is any.
func (c *FakeNodeResourceTopologies) Create(ctx context.Context, nodeResourceTopology *v1alpha1.NodeResourceTopology, opts v1.CreateOptions) (result *v1alpha1.NodeResourceTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(noderesourcetopologiesResource, c.ns, nodeResourceTopology), &v1alpha1.NodeResourceTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeResourceTopology), err
}

// Update takes the representation of a nodeResourceTopology and updates it. Returns the server's representation of the nodeResourceTopology, and an error, if there is any.
func (c *FakeNodeResourceTopologies) Update(ctx context.Context, nodeResourceTopology *v1alpha1.NodeResourceTopology, opts v1.UpdateOptions) (result *v1alpha1.NodeResourceTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(noderesourcetopologiesResource, c.ns, nodeResourceTopology), &v1alpha1.NodeResourceTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeResourceTopology), err
}

// Delete takes name of the nodeResourceTopology and deletes it. Returns an error if one occurs.
func (c *FakeNodeResourceTopologies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(noderesourcetopologiesResource, c.ns, name), &v1alpha1.NodeResourceTopology{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeResourceTopologies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(noderesourcetopologiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeResourceTopologyList{})
	return err
}

// Patch applies the patch and returns the patched nodeResourceTopology.
func (c *FakeNodeResourceTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeResourceTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(noderesourcetopologiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NodeResourceTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeResourceTopology), err
}
