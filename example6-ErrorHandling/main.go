package main

import (
	"errors"
	"fmt"
)

type errUserNameExist struct {
	UserName string
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}

func checkUserNameExist(username string) (bool, error) {
	if username == "foo" {
		return true, fmt.Errorf("username %s already exist", username)
	}
	if username == "bar" {
		return true, errors.New("username bar already exist")
	}

	// customize error
	if username == "foo2" {
		return true, errUserNameExist{UserName: username}
	}
	return false, nil
}

func main() {
	// _, err is check error is exist or not
	if _, err := checkUserNameExist("foo"); err != nil {
		fmt.Println(err)
	}

	if _, err := checkUserNameExist("bar"); err != nil {
		fmt.Println(err)
	}

	if _, err := checkUserNameExist("no"); err != nil {
		fmt.Println(err)
	}

	if _, err := checkUserNameExist("foo2"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		}
	}

	if _, err := checkUserNameExist("foo"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("Error type is not UserNameExist")
		}
	}

}
