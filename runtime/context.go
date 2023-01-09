package runtime

import "context"

var (
	SystemContext       context.Context
	SystemContextCancel context.CancelFunc
)

func SetupContext(parentCh <-chan struct{}) {
	SystemContext, SystemContextCancel = contextForChannel(parentCh)
}

func contextForChannel(parentCh <-chan struct{}) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-parentCh:
			cancel()
		case <-ctx.Done():
		}
	}()
	return ctx, cancel
}
