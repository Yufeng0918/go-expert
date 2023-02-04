package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

func send() {

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	time.Sleep(time.Millisecond * 200)
	values <- value
}
func main() {

	defer close(values)
	go send()
	value := <-values
	fmt.Printf("receive %v\n", value)

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for r := range c {
		fmt.Printf("read: %v\n", r)
	}
	time.Sleep(time.Second * 2)
}
