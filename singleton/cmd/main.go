package main

import (
	"fmt"
	"singleton/pkg/singleton"
)

func main() {
	obj := singleton.NewSingleton("hoge")
	obj1 := singleton.GetSingletonInstance()
	obj2 := singleton.GetSingletonInstance()
	obj3 := singleton.NewSingleton("huga")

	fmt.Println(obj.Title)
	if obj == obj1 {
		if obj1 == obj2 {
			fmt.Println("obj,obj1 and obj2 are the same")
		} else {
			fmt.Println("not Singleton")
		}
	} else {
		fmt.Println("not Singleton")
	}

    fmt.Println(obj3.Title)
	if obj == obj3 {
		fmt.Println("obj and obj3 are the same")
	} else {
		fmt.Println("not Singleton")
	}

}
