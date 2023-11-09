package hvac

func newPowerLed() *powerLed {
	return &powerLed{
		led: make(chan struct{}),
	}
}

type powerLed struct {
	led chan struct{}
}

func (h *powerLed) Running() <-chan struct{} {
	return h.led
}

func (h *powerLed) activate() {
	close(h.led)
}

func (h *powerLed) deactivate() {
	h.led = make(chan struct{})
}
