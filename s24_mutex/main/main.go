package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 0

var lock sync.Mutex

func add() {
	defer lock.Unlock()
	lock.Lock()
	i++
	fmt.Printf("i++ %v\n", i)
}

func sub() {
	defer lock.Unlock()
	lock.Lock()
	i--
	fmt.Printf("i-- %v\n", i)
}

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("final :  %v", i)
}
