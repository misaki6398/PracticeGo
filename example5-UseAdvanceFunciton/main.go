package main

import (
	"fmt"
	"strings"
)

// single return value
func add(i, j int) int {
	return i + j
}

// mutiple return value
func swap(i, j int) (int, int) {
	return j, i
}

// function reutrn
func foo() func() int {
	return func() int {
		return 100
	}
}

func getUserListSQL(username, email string) string { // it's diffcult to extend
	sql := "select * from user"
	where := []string{}
	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}
	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}
	return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
	username string
	email    string
}

func getUserListSQL2(opts searchOpts) string { // it's easy to extend
	sql := "select * from user"
	where := []string{}
	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}
	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	// single return value
	fmt.Println(add(1, 2))

	// mutiple return value
	a := 1
	b := 2
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	a, b = swap(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// or a,b = b,a it's ok

	bar := foo()
	fmt.Printf("type of foo is %T\n", bar)
	fmt.Println(bar())

	bar2 := func(i, j float32) float32 {
		return i + j
	}

	fmt.Printf("type of foo is %T\n", bar2)
	fmt.Println(bar2(1.1, 2.3))

	// anonymous function
	func() {
		fmt.Println("Hello World 3")
	}()

	// will be run on background
	go func(i, j int) {
		fmt.Println(i + j)
	}(1, 2)

	fmt.Println(getUserListSQL("asaki0510", ""))
	fmt.Println(getUserListSQL("asaki0510", "test@gmail.com"))
	fmt.Println(getUserListSQL2(searchOpts{
		username: "asaki0510",
		email:    "test@gmail.com",
	}))

}
