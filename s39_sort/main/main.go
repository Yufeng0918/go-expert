package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("%v %v\n", now.Year(), now.Month())
}
