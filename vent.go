package hvac

import "time"

func NewVent() *Vent {
	return &Vent{
		powerLed: newPowerLed(),
	}
}

type Vent struct {
	*powerLed
}

// Callers can only call Run once on a instance.
func (h *Vent) Run() {
	for {
		h.powerLed.activate()
		h.uploadState()
		time.Sleep(time.Second)
	}
}

func (h *Vent) uploadState() {}
