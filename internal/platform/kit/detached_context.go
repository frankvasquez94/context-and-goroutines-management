package kit

import (
	"context"
	"time"
)

type detachedContext struct {
	parent context.Context
}

// DetachedContext
// Create a detached context holding all parent values, such as telemetry info.
// This something like context Background that is never cancelled, but holds the parent info.
func DetachedContext(ctx context.Context) context.Context {
	return &detachedContext{parent: ctx}
}

func (d *detachedContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false // Never has a deadline
}

func (d *detachedContext) Done() <-chan struct{} {
	return nil // Never done
}

func (d *detachedContext) Err() error {
	return nil // Never has an error
}

func (d *detachedContext) Value(key any) any {
	return d.parent.Value(key) // Preserve values from parent
}
