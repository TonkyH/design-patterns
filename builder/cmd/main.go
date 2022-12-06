package main

import (
	"builder/pkg/builder"
	"builder/pkg/director"
	"fmt"
)

func main() {
	textBuilder := builder.NewTextBuilder()
	director := director.NewDirector(textBuilder)
	director.Construct()
	result := textBuilder.GetTextResult()
	fmt.Println(result)
}
