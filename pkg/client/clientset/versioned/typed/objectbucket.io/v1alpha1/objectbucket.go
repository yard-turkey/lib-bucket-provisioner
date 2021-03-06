/*
Copyright 2019 Red Hat Inc.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kube-object-storage/lib-bucket-provisioner/pkg/apis/objectbucket.io/v1alpha1"
	scheme "github.com/kube-object-storage/lib-bucket-provisioner/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ObjectBucketsGetter has a method to return a ObjectBucketInterface.
// A group's client should implement this interface.
type ObjectBucketsGetter interface {
	ObjectBuckets() ObjectBucketInterface
}

// ObjectBucketInterface has methods to work with ObjectBucket resources.
type ObjectBucketInterface interface {
	Create(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.CreateOptions) (*v1alpha1.ObjectBucket, error)
	Update(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.UpdateOptions) (*v1alpha1.ObjectBucket, error)
	UpdateStatus(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.UpdateOptions) (*v1alpha1.ObjectBucket, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ObjectBucket, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ObjectBucketList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectBucket, err error)
	ObjectBucketExpansion
}

// objectBuckets implements ObjectBucketInterface
type objectBuckets struct {
	client rest.Interface
}

// newObjectBuckets returns a ObjectBuckets
func newObjectBuckets(c *ObjectbucketV1alpha1Client) *objectBuckets {
	return &objectBuckets{
		client: c.RESTClient(),
	}
}

// Get takes name of the objectBucket, and returns the corresponding objectBucket object, and an error if there is any.
func (c *objectBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectBucket, err error) {
	result = &v1alpha1.ObjectBucket{}
	err = c.client.Get().
		Resource("objectbuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ObjectBuckets that match those selectors.
func (c *objectBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ObjectBucketList{}
	err = c.client.Get().
		Resource("objectbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested objectBuckets.
func (c *objectBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("objectbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a objectBucket and creates it.  Returns the server's representation of the objectBucket, and an error, if there is any.
func (c *objectBuckets) Create(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.CreateOptions) (result *v1alpha1.ObjectBucket, err error) {
	result = &v1alpha1.ObjectBucket{}
	err = c.client.Post().
		Resource("objectbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectBucket).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a objectBucket and updates it. Returns the server's representation of the objectBucket, and an error, if there is any.
func (c *objectBuckets) Update(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.UpdateOptions) (result *v1alpha1.ObjectBucket, err error) {
	result = &v1alpha1.ObjectBucket{}
	err = c.client.Put().
		Resource("objectbuckets").
		Name(objectBucket.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectBucket).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *objectBuckets) UpdateStatus(ctx context.Context, objectBucket *v1alpha1.ObjectBucket, opts v1.UpdateOptions) (result *v1alpha1.ObjectBucket, err error) {
	result = &v1alpha1.ObjectBucket{}
	err = c.client.Put().
		Resource("objectbuckets").
		Name(objectBucket.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectBucket).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the objectBucket and deletes it. Returns an error if one occurs.
func (c *objectBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("objectbuckets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *objectBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("objectbuckets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched objectBucket.
func (c *objectBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectBucket, err error) {
	result = &v1alpha1.ObjectBucket{}
	err = c.client.Patch(pt).
		Resource("objectbuckets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
