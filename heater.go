package hvac

import "time"

func NewHeater() *Heater {
	return &Heater{
		running: make(chan struct{}),
	}
}

type Heater struct {
	running chan struct{}
}

func (h *Heater) Run() {
	for {
		time.Sleep(time.Second)
		close(h.running)
	}
}

func (h *Heater) Running() <-chan struct{} {
	return h.running
}
