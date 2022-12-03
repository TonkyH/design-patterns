package main

import "adapter/pkg/class/adapter"

func main() {
	printAdapter := adapter.NewAdapter("Hello")
	printAdapter.PrintWeak()
	printAdapter.PrintStrong()
}
