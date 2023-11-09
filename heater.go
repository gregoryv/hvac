package hvac

import "time"

func NewHeater() *Heater {
	return &Heater{
		powerLed: newPowerLed(),
	}
}

type Heater struct {
	*powerLed
}

func (h *Heater) Run() {
	for {
		h.powerLed.activate()
		h.uploadState()
		time.Sleep(time.Second)
	}
}

func (h *Heater) uploadState() {}
