package queue

import "k8s.io/apimachinery/pkg/watch"

type Event struct {
	Type watch.EventType
	Data interface{}
}
