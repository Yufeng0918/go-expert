package main

import (
	"fmt"
)

func testAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i %v\n", i)
}

func testPanic() {

	defer fmt.Printf("after panic")
	panic("error")
	fmt.Printf("print after panic")
}

func main() {

	i := new(int)
	fmt.Printf("i %T\n", i)
	c := make([]int, 2)
	fmt.Printf("c %T\n", c)

	testAppend()
	testPanic()
	fmt.Println("end")
}
