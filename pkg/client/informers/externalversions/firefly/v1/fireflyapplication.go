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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	firefly_v1 "github.com/frodehus/firefly/pkg/apis/firefly/v1"
	versioned "github.com/frodehus/firefly/pkg/client/clientset/versioned"
	internalinterfaces "github.com/frodehus/firefly/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/frodehus/firefly/pkg/client/listers/firefly/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FireflyApplicationInformer provides access to a shared informer and lister for
// FireflyApplications.
type FireflyApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.FireflyApplicationLister
}

type fireflyApplicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFireflyApplicationInformer constructs a new informer for FireflyApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFireflyApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFireflyApplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFireflyApplicationInformer constructs a new informer for FireflyApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFireflyApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PepperprovesapointV1().FireflyApplications(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PepperprovesapointV1().FireflyApplications(namespace).Watch(options)
			},
		},
		&firefly_v1.FireflyApplication{},
		resyncPeriod,
		indexers,
	)
}

func (f *fireflyApplicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFireflyApplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fireflyApplicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&firefly_v1.FireflyApplication{}, f.defaultInformer)
}

func (f *fireflyApplicationInformer) Lister() v1.FireflyApplicationLister {
	return v1.NewFireflyApplicationLister(f.Informer().GetIndexer())
}