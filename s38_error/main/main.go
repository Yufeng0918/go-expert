package main

import (
	"errors"
	"fmt"
)

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
}
