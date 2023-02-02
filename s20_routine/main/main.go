package main

import (
	"fmt"
	"time"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {

	go show("go")
	go show("java")

	time.Sleep(time.Second * 1)
}
