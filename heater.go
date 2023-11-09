package hvac

import (
	"context"
	"time"
)

func NewHeater() *Heater {
	return &Heater{
		powerLed: newPowerLed(),
	}
}

type Heater struct {
	*powerLed
}

func (h *Heater) Run(ctx context.Context) {
	h.powerLed.activate()
	defer h.powerLed.deactivate()

	for {
		h.uploadState()
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}

func (h *Heater) uploadState() {}
