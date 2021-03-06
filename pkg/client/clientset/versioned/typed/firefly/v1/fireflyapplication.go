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

package v1

import (
	v1 "github.com/frodehus/firefly/pkg/apis/firefly/v1"
	scheme "github.com/frodehus/firefly/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FireflyApplicationsGetter has a method to return a FireflyApplicationInterface.
// A group's client should implement this interface.
type FireflyApplicationsGetter interface {
	FireflyApplications(namespace string) FireflyApplicationInterface
}

// FireflyApplicationInterface has methods to work with FireflyApplication resources.
type FireflyApplicationInterface interface {
	Create(*v1.FireflyApplication) (*v1.FireflyApplication, error)
	Update(*v1.FireflyApplication) (*v1.FireflyApplication, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.FireflyApplication, error)
	List(opts meta_v1.ListOptions) (*v1.FireflyApplicationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.FireflyApplication, err error)
	FireflyApplicationExpansion
}

// fireflyApplications implements FireflyApplicationInterface
type fireflyApplications struct {
	client rest.Interface
	ns     string
}

// newFireflyApplications returns a FireflyApplications
func newFireflyApplications(c *PepperprovesapointV1Client, namespace string) *fireflyApplications {
	return &fireflyApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fireflyApplication, and returns the corresponding fireflyApplication object, and an error if there is any.
func (c *fireflyApplications) Get(name string, options meta_v1.GetOptions) (result *v1.FireflyApplication, err error) {
	result = &v1.FireflyApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fireflyapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FireflyApplications that match those selectors.
func (c *fireflyApplications) List(opts meta_v1.ListOptions) (result *v1.FireflyApplicationList, err error) {
	result = &v1.FireflyApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fireflyapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fireflyApplications.
func (c *fireflyApplications) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fireflyapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a fireflyApplication and creates it.  Returns the server's representation of the fireflyApplication, and an error, if there is any.
func (c *fireflyApplications) Create(fireflyApplication *v1.FireflyApplication) (result *v1.FireflyApplication, err error) {
	result = &v1.FireflyApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fireflyapplications").
		Body(fireflyApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a fireflyApplication and updates it. Returns the server's representation of the fireflyApplication, and an error, if there is any.
func (c *fireflyApplications) Update(fireflyApplication *v1.FireflyApplication) (result *v1.FireflyApplication, err error) {
	result = &v1.FireflyApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fireflyapplications").
		Name(fireflyApplication.Name).
		Body(fireflyApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the fireflyApplication and deletes it. Returns an error if one occurs.
func (c *fireflyApplications) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fireflyapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fireflyApplications) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fireflyapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched fireflyApplication.
func (c *fireflyApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.FireflyApplication, err error) {
	result = &v1.FireflyApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fireflyapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
