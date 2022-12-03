package main

import (
	"factory/pkg/idcard"
)

func main() {
	factory := idcard.NewIdCardFactory()
	card1 := factory.Create("hoge")
	card2 := factory.Create("huga")
	card3 := factory.Create("fiz")

	card1.Use()
	card2.Use()
	card3.Use()

}
