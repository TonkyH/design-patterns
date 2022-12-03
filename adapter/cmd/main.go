package main

import "adapter/pkg/adapter"

func main() {
	printAdapter := adapter.NewAdapter("Hello")
	printAdapter.PrintWeak()
	printAdapter.PrintStrong()
}
