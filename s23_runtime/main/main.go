package main

import (
	"fmt"
	"runtime"
)

func hello(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("msg %v\n", msg)
	}
}

func main() {

	fmt.Printf("cpu %v\n", runtime.NumCPU())
	go hello("java")

	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Printf("i %v\n", i)
	}
	fmt.Println("ending...")
}
