package hvac

import "time"

func NewHeater() *Heater {
	return &Heater{
		PowerLed: NewPowerLed(),
	}
}

type Heater struct {
	*PowerLed
}

func (h *Heater) Run() {
	for {
		h.PowerLed.activate()
		h.uploadState()
		time.Sleep(time.Second)
	}
}

func (h *Heater) uploadState() {}
