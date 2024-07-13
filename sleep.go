// Package timer provides a context-aware Sleep function.
//
// It offers a way to pause execution for a specified duration,
// while respecting context cancellation.
package timer

import (
	"context"
	"time"
)

// Sleep pauses the current goroutine for the specified duration or until the
// context is canceled.
//
// It returns nil if the sleep completes normally, or the context's error if
// the context is canceled before the duration elapses.
func Sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
