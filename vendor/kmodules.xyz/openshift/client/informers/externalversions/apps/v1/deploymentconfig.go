/*
Copyright AppsCode Inc. and Contributors

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

	appsv1 "kmodules.xyz/openshift/apis/apps/v1"
	versioned "kmodules.xyz/openshift/client/clientset/versioned"
	internalinterfaces "kmodules.xyz/openshift/client/informers/externalversions/internalinterfaces"
	v1 "kmodules.xyz/openshift/client/listers/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeploymentConfigInformer provides access to a shared informer and lister for
// DeploymentConfigs.
type DeploymentConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DeploymentConfigLister
}

type deploymentConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeploymentConfigInformer constructs a new informer for DeploymentConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeploymentConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeploymentConfigInformer constructs a new informer for DeploymentConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeploymentConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().DeploymentConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().DeploymentConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&appsv1.DeploymentConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *deploymentConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeploymentConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deploymentConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1.DeploymentConfig{}, f.defaultInformer)
}

func (f *deploymentConfigInformer) Lister() v1.DeploymentConfigLister {
	return v1.NewDeploymentConfigLister(f.Informer().GetIndexer())
}
