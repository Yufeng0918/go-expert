package main

import (
	"fmt"
	"time"
)

var c0 = make(chan int)
var c1 = make(chan string)

func main() {

	go func() {
		defer close(c1)
		defer close(c0)

		c0 <- 100
		c1 <- "hello"
	}()

	for {
		select {
		case r := <-c0:
			fmt.Printf("receive %v\n", r)
		case r := <-c1:
			fmt.Printf("receive %v\n", r)
		}
		time.Sleep(time.Second * 2)
	}
}
