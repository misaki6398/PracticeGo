package main

import (
	"fmt"
)

func checkValue1(s int) {
	switch s {
	case 0:
	case 1:
		fmt.Println("Checked value is ", s)
	}
}

func checkValue2(s int) {
	switch s {
	case 0:
		fallthrough
	case 1:
		fmt.Println("Checked value is ", s)
	}
}

func checkValue3(s int) {
	switch s {
	case 0, 1:
		fmt.Println("Checked value is ", s)
	}
}

func main() {
	// Be careful switch case didn't need break
	fmt.Println("Where is 0?")
	checkValue1(0)
	checkValue1(1)

	// Use fallthrough
	fmt.Println("Use fallthrough")
	checkValue2(0)
	checkValue2(1)

	// Write case together
	fmt.Println("Write case together")
	checkValue3(0)
	checkValue3(1)
}
