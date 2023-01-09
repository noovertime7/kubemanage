package runtime

import (
	"testing"
)

func TestContext(t *testing.T) {
	quit := SetupSignalHandler()
	SetupContext(quit)

	<-SystemContext.Done()
	t.Logf("system exit")
}
