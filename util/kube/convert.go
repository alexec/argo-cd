package kube

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func convertToVersionWithScheme(ctx context.Context, obj *unstructured.Unstructured, group string, version string) (*unstructured.Unstructured, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "convertToVersionWithScheme")
	defer span.Finish()
	s := legacyscheme.Scheme
	object, err := s.ConvertToVersion(obj, runtime.InternalGroupVersioner)
	if err != nil {
		return nil, err
	}
	unmarshalledObj, err := s.ConvertToVersion(object, schema.GroupVersion{Group: group, Version: version})
	if err != nil {
		return nil, err
	}
	unstrBody, err := runtime.DefaultUnstructuredConverter.ToUnstructured(unmarshalledObj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: unstrBody}, nil
}
