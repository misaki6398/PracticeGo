package main

import (
	"fmt"
)

func modify(foo []string) {
	foo[1] = "c"
	fmt.Println("modify foo", foo)
}

func addValue(foo []string) {
	foo = append(foo, "c")
	fmt.Println("modify foo", foo)
}

func addValueReference(foo *[]string) {
	*foo = append(*foo, "c")
	fmt.Println("modify foo", foo)
}

func addValueWithReturn(foo []string) []string {
	foo = append(foo, "c")
	fmt.Println("modify foo", foo)
	return foo
}

func main() {

	foo := []string{"a", "b"}
	fmt.Println("before foo", foo)
	modify(foo) // call by reference
	fmt.Println("after foo", foo)

	foo2 := []string{"a", "b"}
	fmt.Println("before foo2", foo2)
	addValue(foo2) // call by reference
	fmt.Println("after foo2", foo2)

	foo3 := []string{"a", "b"}
	fmt.Println("before foo3", foo3)
	addValueReference(&foo3) // call by reference
	fmt.Println("after foo3", foo3)

	// above methods are not recommand

	// recommand
	foo4 := []string{"a", "b"}
	fmt.Println("before foo4", foo4)
	foo4 = addValueWithReturn(foo4)
	fmt.Println("after foo4", foo4)

	// another trap
	fmt.Println()
	fmt.Println("Another trap")
	foo5 := []string{"a", "b"}
	fmt.Println("before foo", foo5)
	bar := foo5[:1]
	fmt.Println("bar:", bar)
	s1 := append(bar, "c")
	fmt.Println("foo:", foo5)
	fmt.Println("s1:", s1)
	s2 := append(bar, "d")
	fmt.Println("foo:", foo5)
	fmt.Println("s2:", s2)

	// Because of the original variable define is 2 space, append will become 3 space
	s3 := append(bar, "e", "f")
	fmt.Println("foo:", foo5)
	fmt.Println("s3:", s3)

}
