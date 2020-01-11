package testing

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	TestNamespace = "default"
)

func WithAnnotations(annotations map[string]string) func(app *unstructured.Unstructured) {
	return func(app *unstructured.Unstructured) {
		app.SetAnnotations(annotations)
	}
}

func NewApp(name string, modifiers ...func(app *unstructured.Unstructured)) *unstructured.Unstructured {
	app := unstructured.Unstructured{}
	app.SetGroupVersionKind(schema.GroupVersionKind{Group: "argoproj.io", Kind: "application", Version: "v1alpha1"})
	app.SetName(name)
	app.SetNamespace(TestNamespace)
	for i := range modifiers {
		modifiers[i](&app)
	}
	return &app
}
