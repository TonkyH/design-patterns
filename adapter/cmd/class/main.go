package main

import "adapter/pkg/class/adapter"

func main() {
	printAdabter := adapter.NewAdapter("Hello")
	printAdabter.PrintWeak()
	printAdabter.PrintStrong()
}
