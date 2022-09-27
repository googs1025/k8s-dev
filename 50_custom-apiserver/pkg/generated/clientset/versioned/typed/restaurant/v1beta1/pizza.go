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

package v1beta1

import (
	v1beta1 "50_custom-apiserver/pkg/apis/restaurant/v1beta1"
	scheme "50_custom-apiserver/pkg/generated/clientset/versioned/scheme"
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PizzasGetter has a method to return a PizzaInterface.
// A group's client should implement this interface.
type PizzasGetter interface {
	Pizzas(namespace string) PizzaInterface
}

// PizzaInterface has methods to work with Pizza resources.
type PizzaInterface interface {
	Create(ctx context.Context, pizza *v1beta1.Pizza, opts v1.CreateOptions) (*v1beta1.Pizza, error)
	Update(ctx context.Context, pizza *v1beta1.Pizza, opts v1.UpdateOptions) (*v1beta1.Pizza, error)
	UpdateStatus(ctx context.Context, pizza *v1beta1.Pizza, opts v1.UpdateOptions) (*v1beta1.Pizza, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Pizza, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PizzaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Pizza, err error)
	PizzaExpansion
}

// pizzas implements PizzaInterface
type pizzas struct {
	client rest.Interface
	ns     string
}

// newPizzas returns a Pizzas
func newPizzas(c *RestaurantV1beta1Client, namespace string) *pizzas {
	return &pizzas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pizza, and returns the corresponding pizza object, and an error if there is any.
func (c *pizzas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Pizza, err error) {
	result = &v1beta1.Pizza{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pizzas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Pizzas that match those selectors.
func (c *pizzas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PizzaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PizzaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pizzas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pizzas.
func (c *pizzas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pizzas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pizza and creates it.  Returns the server's representation of the pizza, and an error, if there is any.
func (c *pizzas) Create(ctx context.Context, pizza *v1beta1.Pizza, opts v1.CreateOptions) (result *v1beta1.Pizza, err error) {
	result = &v1beta1.Pizza{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pizzas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pizza).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pizza and updates it. Returns the server's representation of the pizza, and an error, if there is any.
func (c *pizzas) Update(ctx context.Context, pizza *v1beta1.Pizza, opts v1.UpdateOptions) (result *v1beta1.Pizza, err error) {
	result = &v1beta1.Pizza{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pizzas").
		Name(pizza.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pizza).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *pizzas) UpdateStatus(ctx context.Context, pizza *v1beta1.Pizza, opts v1.UpdateOptions) (result *v1beta1.Pizza, err error) {
	result = &v1beta1.Pizza{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pizzas").
		Name(pizza.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pizza).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pizza and deletes it. Returns an error if one occurs.
func (c *pizzas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pizzas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pizzas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pizzas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pizza.
func (c *pizzas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Pizza, err error) {
	result = &v1beta1.Pizza{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pizzas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
