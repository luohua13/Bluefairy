package informer

import (
	"Bluefairy/pkg/event"
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
)

func NewEndpointsInformer(endpointsClient v1.EndpointsInterface) cache.Controller {
	return NewInformer(
		"watch-endpoints",
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				options.FieldSelector = fields.Everything().String()
				return endpointsClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				options.Watch = true
				options.FieldSelector = fields.Everything().String()
				return endpointsClient.Watch(context.TODO(), options)
			}},
		&corev1.Endpoints{},
		0,
		ResourceEventHandlerFuncs{
			ResourceEventHandlerFuncs: cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {
					event.AddEndpoints(obj)
				},
				UpdateFunc: func(_, newObj interface{}) {
					event.UpdateEndpoints(newObj)
				},
				DeleteFunc: func(obj interface{}) {
					event.DeleteEndpoints(obj)
				},
			},
			ReplaceFunc: func(objs []interface{}) {
				event.ReplaceEndpoints(objs)
			},
		})
}
