package hvac

import "time"

func NewVent() *Vent {
	return &Vent{
		PowerLed: NewPowerLed(),
	}
}

type Vent struct {
	*PowerLed
}

// Callers can only call Run once on a instance.
func (h *Vent) Run() {
	for {
		h.PowerLed.activate()
		h.uploadState()
		time.Sleep(time.Second)
	}
}

func (h *Vent) uploadState() {}
