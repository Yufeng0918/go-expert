package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Printf("pi %v\n", math.Pi)
	rand.Seed(1)
	fmt.Printf("rand %v\n", rand.Int())
	fmt.Printf("rand %v\n", rand.Int())
}
