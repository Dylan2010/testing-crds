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

package v1alpha1

import (
	time "time"

	catv1alpha1 "github.com/bobcatfish/testing-crds/client-go/pkg/apis/cat/v1alpha1"
	versioned "github.com/bobcatfish/testing-crds/client-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/bobcatfish/testing-crds/client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bobcatfish/testing-crds/client-go/pkg/client/listers/cat/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CatInformer provides access to a shared informer and lister for
// Cats.
type CatInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CatLister
}

type catInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCatInformer constructs a new informer for Cat type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCatInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCatInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCatInformer constructs a new informer for Cat type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCatInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatV1alpha1().Cats(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatV1alpha1().Cats(namespace).Watch(options)
			},
		},
		&catv1alpha1.Cat{},
		resyncPeriod,
		indexers,
	)
}

func (f *catInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCatInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *catInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&catv1alpha1.Cat{}, f.defaultInformer)
}

func (f *catInformer) Lister() v1alpha1.CatLister {
	return v1alpha1.NewCatLister(f.Informer().GetIndexer())
}
