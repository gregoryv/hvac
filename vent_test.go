package hvac

import (
	"testing"
	"time"
)

func TestVent_Run(t *testing.T) {
	h := NewVent()

	go h.Run()

	/*
	   Test routine               Run routine
	                 chan
	           <-  ========
	*/
	select {
	case <-h.Running():
		// ok
		t.Log("ok")

	case <-time.After(10 * time.Millisecond):
		t.Error("too slow")
	}
}
