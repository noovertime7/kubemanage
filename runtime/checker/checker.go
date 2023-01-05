package checker

import "github.com/noovertime7/kubemanage/runtime/queue"

type Checker interface {
	Run()
	Check(event *queue.Event) error
	HandlerErr(err error)
}
