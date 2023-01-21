package wait

import (
	"fmt"
	"testing"
	"time"

	"github.com/noovertime7/kubemanage/runtime"
)

func TestBackoffUntil(t *testing.T) {
	handler := NewDefaultBackoff(3 * time.Second)
	BackoffUntil(func() {
		t.Log("called")
	}, handler, true, runtime.SetupSignalHandler())
}

func TestPollImmediateUntil(t *testing.T) {
	stop := make(chan struct{})
	err := PollImmediateUntil(time.Second*3, func() (done bool, err error) {
		t.Logf("pulled")
		return false, fmt.Errorf("test err")
	}, stop)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("quit")
}
