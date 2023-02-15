package main

import (
	"fmt"
)

// decalre method 1
var foo string
var bar int

// decalre method 2
var (
	foo2 string = "Hello2"
	bar2 int    = 200
)

// decalre const 1
const (
	Monday1  = 1
	Tuesday1 = 2
)

// decalre const 2
const (
	Monday2 = iota + 1
	Tuesday2
	Wednesday2
)

func main() {
	foo = "Hello"
	bar = 100

	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(foo2)
	fmt.Println(bar2)

	// decalre method 3
	foo3 := "Hello3"
	bar3 := 300
	fmt.Println(foo3)
	fmt.Println(bar3)

	fmt.Println("const")
	fmt.Println(Monday1)
	fmt.Println(Tuesday1)

	fmt.Println("iota const")
	fmt.Println(Monday2)
	fmt.Println(Tuesday2)
	fmt.Println(Wednesday2)

}
