/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	"github.com/jetstack-experimental/navigator/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type NavigatorInterface interface {
	RESTClient() rest.Interface
	ElasticsearchClustersGetter
}

// NavigatorClient is used to interact with features provided by the navigator.jetstack.io group.
type NavigatorClient struct {
	restClient rest.Interface
}

func (c *NavigatorClient) ElasticsearchClusters(namespace string) ElasticsearchClusterInterface {
	return newElasticsearchClusters(c, namespace)
}

// NewForConfig creates a new NavigatorClient for the given config.
func NewForConfig(c *rest.Config) (*NavigatorClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NavigatorClient{client}, nil
}

// NewForConfigOrDie creates a new NavigatorClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NavigatorClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NavigatorClient for the given RESTClient.
func New(c rest.Interface) *NavigatorClient {
	return &NavigatorClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("navigator.jetstack.io")
	if err != nil {
		return err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != g.GroupVersion.Group {
		gv := g.GroupVersion
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NavigatorClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}