package main

import (
	prototype "Prototype/internal"
	"fmt"
)

func main() {
	product := prototype.NewConcretePrototype("New Product 1")
	cloneProduct := product.Clone()
	fmt.Println(cloneProduct.GetName())
}
