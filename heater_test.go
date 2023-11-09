package hvac

import (
	"testing"
	"time"
)

func TestHeater_Run(t *testing.T) {
	h := NewHeater()

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
