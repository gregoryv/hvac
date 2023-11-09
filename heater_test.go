package hvac

import (
	"testing"
)

func TestHeater_Run(t *testing.T) {
	h := NewHeater()

	go h.Run()

	/*
			   Test routine               Run routine
		                         chan
		                   <-  ========
	*/
	<-h.Running()
}
