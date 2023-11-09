package hvac

import "time"

func NewVent() *Vent {
	return &Vent{
		running: make(chan struct{}),
	}
}

type Vent struct {
	running chan struct{}
}

// Callers can only call Run once on a instance.
func (h *Vent) Run() {
	for {
		close(h.running)
		h.uploadState()
		time.Sleep(time.Second)
	}
}

func (h *Vent) Running() <-chan struct{} {
	return h.running
}

func (h *Vent) uploadState() {

}
