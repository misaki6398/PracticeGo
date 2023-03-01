package main

import (
	"fmt"
)

type car struct {
	name string
}

func (c car) SetName01(s string) {
	fmt.Printf("1. address is %p", &c)
	c.name = s
}

func (c *car) SetName02(s string) {
	fmt.Printf("2. address is %p", c)
	c.name = s
}

func main() {
	c := &car{}
	fmt.Printf("c address: %p\n", c)
	// pass point
	c.SetName02("bar")
	fmt.Println(c.name)

	// copy stracture
	c.SetName01("foo")
	fmt.Println(c.name)
}
