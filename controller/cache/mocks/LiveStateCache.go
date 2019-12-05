// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	kube "github.com/argoproj/argo-cd/util/kube"
	mock "github.com/stretchr/testify/mock"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
)

// LiveStateCache is an autogenerated mock type for the LiveStateCache type
type LiveStateCache struct {
	mock.Mock
}

// GetManagedLiveObjs provides a mock function with given fields: ctx, a, targetObjs
func (_m *LiveStateCache) GetManagedLiveObjs(ctx context.Context, a *v1alpha1.Application, targetObjs []*unstructured.Unstructured) (map[kube.ResourceKey]*unstructured.Unstructured, error) {
	ret := _m.Called(ctx, a, targetObjs)

	var r0 map[kube.ResourceKey]*unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Application, []*unstructured.Unstructured) map[kube.ResourceKey]*unstructured.Unstructured); ok {
		r0 = rf(ctx, a, targetObjs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[kube.ResourceKey]*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.Application, []*unstructured.Unstructured) error); ok {
		r1 = rf(ctx, a, targetObjs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespaceTopLevelResources provides a mock function with given fields: ctx, server, namespace
func (_m *LiveStateCache) GetNamespaceTopLevelResources(ctx context.Context, server string, namespace string) (map[kube.ResourceKey]v1alpha1.ResourceNode, error) {
	ret := _m.Called(ctx, server, namespace)

	var r0 map[kube.ResourceKey]v1alpha1.ResourceNode
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[kube.ResourceKey]v1alpha1.ResourceNode); ok {
		r0 = rf(ctx, server, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[kube.ResourceKey]v1alpha1.ResourceNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, server, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerVersion provides a mock function with given fields: server
func (_m *LiveStateCache) GetServerVersion(server string) (string, error) {
	ret := _m.Called(server)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(server)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Invalidate provides a mock function with given fields:
func (_m *LiveStateCache) Invalidate() {
	_m.Called()
}

// IsNamespaced provides a mock function with given fields: ctx, server, gk
func (_m *LiveStateCache) IsNamespaced(ctx context.Context, server string, gk schema.GroupKind) (bool, error) {
	ret := _m.Called(ctx, server, gk)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, schema.GroupKind) bool); ok {
		r0 = rf(ctx, server, gk)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, schema.GroupKind) error); ok {
		r1 = rf(ctx, server, gk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IterateHierarchy provides a mock function with given fields: ctx, server, key, action
func (_m *LiveStateCache) IterateHierarchy(ctx context.Context, server string, key kube.ResourceKey, action func(v1alpha1.ResourceNode, string)) error {
	ret := _m.Called(ctx, server, key, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, kube.ResourceKey, func(v1alpha1.ResourceNode, string)) error); ok {
		r0 = rf(ctx, server, key, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *LiveStateCache) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
