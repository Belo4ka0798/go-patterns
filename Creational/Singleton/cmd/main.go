package main

import (
	singleton "Singleton/internal"
	"fmt"
)

func main() {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 != instance2 {
		fmt.Println("Objects are not equal!\n")
	} else {
		fmt.Println("Object are equal!\n")
	}
}
