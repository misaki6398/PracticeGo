package main

import (
	"fmt"

	_ "github.com/go-practice/foo"
)

func init() {
	fmt.Println("Init function")
}

func main() {
	fmt.Println("Main function")

}
