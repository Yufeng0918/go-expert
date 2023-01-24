package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "hello world"
	var s1 = "hello world"
	s3 := "hello world"

	s4 := `
	line1
	line2
	`

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s3)
	fmt.Printf("%v\n", s4)

	msg := strings.Join([]string{s, s1}, ",")
	fmt.Printf("%v\n", msg)

	s2 := "c:\\program files\\"
	fmt.Printf("%v\n", s2)

	s5 := `\\program file`
	fmt.Printf("%v\n", s5)
}
