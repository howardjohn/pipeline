/*
Copyright 2020 The Tekton Authors

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

	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/resolution/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResolutionRequests implements ResolutionRequestInterface
type FakeResolutionRequests struct {
	Fake *FakeResolutionV1alpha1
	ns   string
}

var resolutionrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("resolutionrequests")

var resolutionrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("ResolutionRequest")

// Get takes name of the resolutionRequest, and returns the corresponding resolutionRequest object, and an error if there is any.
func (c *FakeResolutionRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResolutionRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolutionrequestsResource, c.ns, name), &v1alpha1.ResolutionRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolutionRequest), err
}

// List takes label and field selectors, and returns the list of ResolutionRequests that match those selectors.
func (c *FakeResolutionRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolutionRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolutionrequestsResource, resolutionrequestsKind, c.ns, opts), &v1alpha1.ResolutionRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolutionRequestList{ListMeta: obj.(*v1alpha1.ResolutionRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolutionRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolutionRequests.
func (c *FakeResolutionRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolutionrequestsResource, c.ns, opts))

}

// Create takes the representation of a resolutionRequest and creates it.  Returns the server's representation of the resolutionRequest, and an error, if there is any.
func (c *FakeResolutionRequests) Create(ctx context.Context, resolutionRequest *v1alpha1.ResolutionRequest, opts v1.CreateOptions) (result *v1alpha1.ResolutionRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolutionrequestsResource, c.ns, resolutionRequest), &v1alpha1.ResolutionRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolutionRequest), err
}

// Update takes the representation of a resolutionRequest and updates it. Returns the server's representation of the resolutionRequest, and an error, if there is any.
func (c *FakeResolutionRequests) Update(ctx context.Context, resolutionRequest *v1alpha1.ResolutionRequest, opts v1.UpdateOptions) (result *v1alpha1.ResolutionRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolutionrequestsResource, c.ns, resolutionRequest), &v1alpha1.ResolutionRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolutionRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolutionRequests) UpdateStatus(ctx context.Context, resolutionRequest *v1alpha1.ResolutionRequest, opts v1.UpdateOptions) (*v1alpha1.ResolutionRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolutionrequestsResource, "status", c.ns, resolutionRequest), &v1alpha1.ResolutionRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolutionRequest), err
}

// Delete takes name of the resolutionRequest and deletes it. Returns an error if one occurs.
func (c *FakeResolutionRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resolutionrequestsResource, c.ns, name, opts), &v1alpha1.ResolutionRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolutionRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolutionrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolutionRequestList{})
	return err
}

// Patch applies the patch and returns the patched resolutionRequest.
func (c *FakeResolutionRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResolutionRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolutionrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResolutionRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolutionRequest), err
}
