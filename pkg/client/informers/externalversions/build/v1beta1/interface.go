// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/shipwright-io/build/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// BuildRuns returns a BuildRunInformer.
	BuildRuns() BuildRunInformer
	// BuildStrategies returns a BuildStrategyInformer.
	BuildStrategies() BuildStrategyInformer
	// ClusterBuildStrategies returns a ClusterBuildStrategyInformer.
	ClusterBuildStrategies() ClusterBuildStrategyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BuildRuns returns a BuildRunInformer.
func (v *version) BuildRuns() BuildRunInformer {
	return &buildRunInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BuildStrategies returns a BuildStrategyInformer.
func (v *version) BuildStrategies() BuildStrategyInformer {
	return &buildStrategyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterBuildStrategies returns a ClusterBuildStrategyInformer.
func (v *version) ClusterBuildStrategies() ClusterBuildStrategyInformer {
	return &clusterBuildStrategyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
