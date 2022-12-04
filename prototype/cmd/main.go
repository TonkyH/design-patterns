package main

import (
	"prototype/pkg/client"
	"prototype/pkg/concrete-prototype"
)

func main() {
	manager := client.NewManager()
	mbox := concreteprototype.NewMessageBox("*")

	manager.Register("warning box", mbox)

	p1 := manager.Create("warning box")
	p1.Use("Hello World")
}
