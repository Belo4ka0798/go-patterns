package main

import (
	builder "Builder/internal"
	"fmt"
)

func main() {
	// some code
	p := builder.Product{}
	cb := builder.ConcreteBuilder{&p}
	dir := builder.Director{&cb}
	dir.Construct()
	fmt.Println(p.Show())
}
