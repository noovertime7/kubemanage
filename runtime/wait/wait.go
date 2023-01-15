package wait

import (
	"context"
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/utils/clock"
)

type BackoffManager interface {
	Backoff() clock.Timer
}

type defaultBackoffHandler struct {
	duration time.Duration
	Clock    clock.Clock
}

func NewDefaultBackoff(duration time.Duration) BackoffManager {
	return &defaultBackoffHandler{duration: duration, Clock: clock.RealClock{}}
}

func (t *defaultBackoffHandler) Backoff() clock.Timer {
	return t.Clock.NewTimer(t.duration)
}

func BackoffUntil(f func(), backoff BackoffManager, sliding bool, stopCh <-chan struct{}) {
	var t clock.Timer
	for {
		select {
		case <-stopCh:
			return
		default:
		}

		if !sliding {
			t = backoff.Backoff()
		}

		func() {
			defer runtime.HandleCrash()
			f()
		}()

		if sliding {
			t = backoff.Backoff()
		}

		// NOTE: b/c there is no priority selection in golang
		// it is possible for this to race, meaning we could
		// trigger t.C and stopCh, and t.C select falls through.
		// In order to mitigate we re-check stopCh at the beginning
		// of every loop to prevent extra executions of f().
		select {
		case <-stopCh:
			if !t.Stop() {
				<-t.C()
			}
			return
		case <-t.C():
		}
	}
}

type ConditionFunc func() (done bool, err error)

// WithContext converts a ConditionFunc into a ConditionWithContextFunc
func (cf ConditionFunc) WithContext() ConditionWithContextFunc {
	return func(context.Context) (done bool, err error) {
		return cf()
	}
}

// ConditionWithContextFunc returns true if the condition is satisfied, or an error
// if the loop should be aborted.
//
// The caller passes along a context that can be used by the condition function.
type ConditionWithContextFunc func(context.Context) (done bool, err error)
type WaitWithContextFunc func(ctx context.Context) <-chan struct{}

func PollImmediateUntil(interval time.Duration, condition ConditionFunc, stopCh <-chan struct{}) error {
	ctx, cancel := ContextForChannel(stopCh)
	defer cancel()
	return PollImmediateUntilWithContext(ctx, interval, condition.WithContext())
}

func PollImmediateUntilWithContext(ctx context.Context, interval time.Duration, condition ConditionWithContextFunc) error {
	return poll(ctx, true, poller(interval, 0), condition)
}

func poll(ctx context.Context, immediate bool, wait WaitWithContextFunc, condition ConditionWithContextFunc) error {
	if immediate {
		done, err := runConditionWithCrashProtectionWithContext(ctx, condition)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}

	select {
	case <-ctx.Done():
		// returning ctx.Err() will break backward compatibility
		return errors.New("timed out waiting for the condition")
	default:
		return WaitForWithContext(ctx, wait, condition)
	}
}

func WaitForWithContext(ctx context.Context, wait WaitWithContextFunc, fn ConditionWithContextFunc) error {
	waitCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := wait(waitCtx)
	for {
		select {
		case _, open := <-c:
			ok, err := runConditionWithCrashProtectionWithContext(ctx, fn)
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
			if !open {
				return errors.New("timed out waiting for the condition")
			}
		case <-ctx.Done():
			// returning ctx.Err() will break backward compatibility
			return errors.New("timed out waiting for the condition")
		}
	}
}

func poller(interval, timeout time.Duration) WaitWithContextFunc {
	return WaitWithContextFunc(func(ctx context.Context) <-chan struct{} {
		ch := make(chan struct{})

		go func() {
			defer close(ch)

			tick := time.NewTicker(interval)
			defer tick.Stop()

			var after <-chan time.Time
			if timeout != 0 {
				// time.After is more convenient, but it
				// potentially leaves timers around much longer
				// than necessary if we exit early.
				timer := time.NewTimer(timeout)
				after = timer.C
				defer timer.Stop()
			}

			for {
				select {
				case <-tick.C:
					// If the consumer isn't ready for this signal drop it and
					// check the other channels.
					select {
					case ch <- struct{}{}:
					default:
					}
				case <-after:
					return
				case <-ctx.Done():
					return
				}
			}
		}()

		return ch
	})
}

func ContextForChannel(parentCh <-chan struct{}) (context.Context, context.CancelFunc) {
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

// runConditionWithCrashProtectionWithContext runs a
// ConditionWithContextFunc with crash protection.
func runConditionWithCrashProtectionWithContext(ctx context.Context, condition ConditionWithContextFunc) (bool, error) {
	defer runtime.HandleCrash()
	return condition(ctx)
}
