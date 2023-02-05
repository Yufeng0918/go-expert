package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int32 = 0

var lock sync.Mutex

func add() {
	atomic.AddInt32(&i, 1)
	fmt.Printf("read i %v\n", atomic.LoadInt32(&i))
}

func sub() {
	atomic.AddInt32(&i, -1)
	fmt.Printf("read i %v\n", atomic.LoadInt32(&i))
}

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("final :  %v", i)
}
