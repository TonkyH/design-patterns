package abstraction

import "bridge/pkg/implementor"

type Display struct {
	Impl implementor.DisplayImpl
}

func NewDisplay(impl implementor.DisplayImpl) *Display {
	return &Display{Impl: impl}
}

func (d *Display) Open() {
	d.Impl.RawOpen()
}

func (d *Display) PrindBody() {
	d.Impl.RawPrint()
}

func (d *Display) Close() {
	d.Impl.RawClose()
}

func (d *Display) Display() {
	d.Open()
	d.PrindBody()
	d.Close()
}
