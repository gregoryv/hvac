package hvac

import (
	"context"
	"testing"
	"time"
)

func TestHeater_RunCanBeStopped(t *testing.T) {
	h := NewHeater()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go cancel()

	h.Run(ctx)
	t.Log("stopped")
}

func TestHeater_Run(t *testing.T) {
	h := NewHeater()

	ctx := context.Background()
	go h.Run(ctx)

	select {
	case <-h.Running():
		// ok
		t.Log("ok")
	case <-time.After(10 * time.Millisecond):
		t.Error("too slow")
	}
}
