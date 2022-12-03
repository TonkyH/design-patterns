package adapter

import (
	"adapter/pkg/banner"
	"adapter/pkg/print"
)

type Adapter struct {
	Banner *banner.Banner
}

func NewAdapter(text string) print.Print {
	banner := banner.NewBanner(text)
	return &Adapter{
		Banner: banner,
	}
}

func (a *Adapter) PrintWeak() {
	a.Banner.ShowWithPaten()
}

func (a *Adapter) PrintStrong() {
	a.Banner.ShowWithSter()
}
