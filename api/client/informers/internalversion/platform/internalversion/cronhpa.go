/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientsetinternalversion "tkestack.io/tke/api/client/clientset/internalversion"
	internalinterfaces "tkestack.io/tke/api/client/informers/internalversion/internalinterfaces"
	internalversion "tkestack.io/tke/api/client/listers/platform/internalversion"
	platform "tkestack.io/tke/api/platform"
)

// CronHPAInformer provides access to a shared informer and lister for
// CronHPAs.
type CronHPAInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CronHPALister
}

type cronHPAInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCronHPAInformer constructs a new informer for CronHPA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronHPAInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronHPAInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCronHPAInformer constructs a new informer for CronHPA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronHPAInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Platform().CronHPAs().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Platform().CronHPAs().Watch(options)
			},
		},
		&platform.CronHPA{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronHPAInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronHPAInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronHPAInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&platform.CronHPA{}, f.defaultInformer)
}

func (f *cronHPAInformer) Lister() internalversion.CronHPALister {
	return internalversion.NewCronHPALister(f.Informer().GetIndexer())
}
