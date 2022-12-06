package director

import (
	"builder/pkg/builder"
)

type Director struct {
	builder builder.IBuilder
}

func NewDirector(builder builder.IBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("一般的な挨拶")
	d.builder.MakeItems([]string{
		"How are you?",
		"Hello",
		"Hi.",
	})

	d.builder.MakeString("時間帯に応じた挨拶")
	d.builder.MakeItems([]string{
		"Good morning",
		"Good afternoon",
		"Good evening",
	})

	d.builder.Close()
}
