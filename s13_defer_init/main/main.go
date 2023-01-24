package main

import "fmt"

var i int = initVar()

func init() {
	fmt.Println("init")
}

func initVar() int {

	fmt.Println("init var")
	return 0
}

func main() {

	fmt.Println("main")
	defer fmt.Println("hello go")
	defer fmt.Println("hello java")
	defer fmt.Println("hello python")
}
