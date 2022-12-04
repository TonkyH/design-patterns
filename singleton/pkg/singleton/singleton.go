package singleton

import (
	"fmt"
)

type Singleton struct {
	Title string
}

var singletonInstance *Singleton

func NewSingleton(title string) *Singleton {
	if singletonInstance == nil {
		singletonInstance = &Singleton{Title: title}
	} else {
		fmt.Println("singletonInstance is already exist")
	}
	return singletonInstance
}

func GetSingletonInstance() *Singleton {
	return singletonInstance
}
