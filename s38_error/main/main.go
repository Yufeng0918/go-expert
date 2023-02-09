package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func check(s string) (string, error) {

	if s == "" {
		err := errors.New("input is empty")
		return "", err
	} else {
		return s, nil
	}

}

func main() {

	s, err := check("")
	fmt.Println("%v %v\n", s, err)
	s, err = check("a")
	fmt.Println("%v %v\n", s, err)
	fmt.Println("%v\n", MyError{time.Now(), "file system is empty"})
}
