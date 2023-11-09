package hvac

func NewPowerLed() *PowerLed {
	return &PowerLed{
		led: make(chan struct{}),
	}
}

type PowerLed struct {
	led chan struct{}
}

func (h *PowerLed) Running() <-chan struct{} {
	return h.led
}

func (h *PowerLed) activate() {
	close(h.led)
}
