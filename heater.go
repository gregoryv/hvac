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
		close(h.running)
		time.Sleep(time.Second)
	}
}

func (h *Heater) Running() <-chan struct{} {
	return h.running
}
