// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/paipaoso/pytorch-operator/pkg/apis/pytorch/v1"
	scheme "github.com/paipaoso/pytorch-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PyTorchJobsGetter has a method to return a PyTorchJobInterface.
// A group's client should implement this interface.
type PyTorchJobsGetter interface {
	PyTorchJobs(namespace string) PyTorchJobInterface
}

// PyTorchJobInterface has methods to work with PyTorchJob resources.
type PyTorchJobInterface interface {
	Create(*v1.PyTorchJob) (*v1.PyTorchJob, error)
	Update(*v1.PyTorchJob) (*v1.PyTorchJob, error)
	UpdateStatus(*v1.PyTorchJob) (*v1.PyTorchJob, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PyTorchJob, error)
	List(opts metav1.ListOptions) (*v1.PyTorchJobList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PyTorchJob, err error)
	PyTorchJobExpansion
}

// pyTorchJobs implements PyTorchJobInterface
type pyTorchJobs struct {
	client rest.Interface
	ns     string
}

// newPyTorchJobs returns a PyTorchJobs
func newPyTorchJobs(c *KubeflowV1Client, namespace string) *pyTorchJobs {
	return &pyTorchJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pyTorchJob, and returns the corresponding pyTorchJob object, and an error if there is any.
func (c *pyTorchJobs) Get(name string, options metav1.GetOptions) (result *v1.PyTorchJob, err error) {
	result = &v1.PyTorchJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pytorchjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PyTorchJobs that match those selectors.
func (c *pyTorchJobs) List(opts metav1.ListOptions) (result *v1.PyTorchJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PyTorchJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pytorchjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pyTorchJobs.
func (c *pyTorchJobs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pytorchjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pyTorchJob and creates it.  Returns the server's representation of the pyTorchJob, and an error, if there is any.
func (c *pyTorchJobs) Create(pyTorchJob *v1.PyTorchJob) (result *v1.PyTorchJob, err error) {
	result = &v1.PyTorchJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pytorchjobs").
		Body(pyTorchJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pyTorchJob and updates it. Returns the server's representation of the pyTorchJob, and an error, if there is any.
func (c *pyTorchJobs) Update(pyTorchJob *v1.PyTorchJob) (result *v1.PyTorchJob, err error) {
	result = &v1.PyTorchJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pytorchjobs").
		Name(pyTorchJob.Name).
		Body(pyTorchJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pyTorchJobs) UpdateStatus(pyTorchJob *v1.PyTorchJob) (result *v1.PyTorchJob, err error) {
	result = &v1.PyTorchJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pytorchjobs").
		Name(pyTorchJob.Name).
		SubResource("status").
		Body(pyTorchJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the pyTorchJob and deletes it. Returns an error if one occurs.
func (c *pyTorchJobs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pytorchjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pyTorchJobs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pytorchjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pyTorchJob.
func (c *pyTorchJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PyTorchJob, err error) {
	result = &v1.PyTorchJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pytorchjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
