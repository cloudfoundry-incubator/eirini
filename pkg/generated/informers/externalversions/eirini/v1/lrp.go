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
	"context"
	time "time"

	eiriniv1 "code.cloudfoundry.org/eirini/pkg/apis/eirini/v1"
	versioned "code.cloudfoundry.org/eirini/pkg/generated/clientset/versioned"
	internalinterfaces "code.cloudfoundry.org/eirini/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "code.cloudfoundry.org/eirini/pkg/generated/listers/eirini/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LRPInformer provides access to a shared informer and lister for
// LRPs.
type LRPInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.LRPLister
}

type lRPInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLRPInformer constructs a new informer for LRP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLRPInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLRPInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLRPInformer constructs a new informer for LRP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLRPInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EiriniV1().LRPs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EiriniV1().LRPs(namespace).Watch(context.TODO(), options)
			},
		},
		&eiriniv1.LRP{},
		resyncPeriod,
		indexers,
	)
}

func (f *lRPInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLRPInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lRPInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eiriniv1.LRP{}, f.defaultInformer)
}

func (f *lRPInformer) Lister() v1.LRPLister {
	return v1.NewLRPLister(f.Informer().GetIndexer())
}
