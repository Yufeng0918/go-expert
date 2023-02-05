package main

import (
	"fmt"
	"time"
)

var c0 = make(chan int)
var c1 = make(chan string)

func main() {

	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time now: %v\n", time.Now())

	t1 := <-timer.C
	fmt.Printf("t1: %v\n", t1)

	<-time.After(time.Second * 2)
	fmt.Printf("sleep after 2 sec\n")

	fmt.Printf("start new timer\n")
	timer0 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer0.C
		fmt.Printf("go function\n")
	}()

	timer0.Stop()
	fmt.Printf("stomp timer 0\n")
}
