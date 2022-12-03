package adapter

import (
	"adapter/pkg/adaptee"
	"adapter/pkg/target"
)

type Adapter struct {
	Adaptee *adaptee.Adaptee
}

func NewAdapter(text string) target.Target {
	adaptee := adaptee.NewAdaptee(text)
	return &Adapter{
		Adaptee: adaptee,
	}
}

func (a *Adapter) PrintWeak() {
	a.Adaptee.ShowWithPaten()
}

func (a *Adapter) PrintStrong() {
	a.Adaptee.ShowWithSter()
}
