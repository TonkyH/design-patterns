package main

import (
	"bridge/pkg/abstraction"
	"bridge/pkg/implementor"
)

func main() {
	d1 := abstraction.NewDisplay(
		implementor.NewStringDisplay("Hello Display"),
	)

	d2 := abstraction.NewCountDisplay(
		implementor.NewStringDisplay("Hello CountDisplay"),
	)

	d3 := abstraction.NewCountDisplay(
		implementor.NewStringDisplay("Hello CountDisplay2"),
	)

	d1.Display()
	d2.Display.Display()
	d3.Display.Display()
	d3.MultiDisplay(5)
}
