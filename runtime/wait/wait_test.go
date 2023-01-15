package wait

import (
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
