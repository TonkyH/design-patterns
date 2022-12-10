package abstraction

import "bridge/pkg/implementor"

type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl implementor.DisplayImpl) *CountDisplay {
	return &CountDisplay{
		Display: &Display{
			Impl: impl,
		},
	}
}

func (d *CountDisplay) MultiDisplay(times int) {
	d.Open()
	for i := 0; i < times; i++ {
		d.PrindBody()
	}
	d.Close()
}
