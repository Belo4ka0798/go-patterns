package main

import (
	facm "FactoryMethod/internal"
	"fmt"
)

func main() {
	var a facm.ConcreteCreator
	fmt.Println(a.GetProduct("A").Use())
}
