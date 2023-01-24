package main

import (
	"fmt"
	"math"
	"unsafe"
)

func c() {

}

func main() {

	var a = [3]int{1, 2, 3}
	fmt.Printf("%T\n", a)

	var b = []int{4, 5, 6}
	fmt.Printf("%T\n", b)

	fmt.Printf("%T\n", c)

	var d int32 = 5
	fmt.Printf("%T %v, %v\n", unsafe.Sizeof(d), math.MinInt32, math.MaxInt32)

	var e int32 = 126
	fmt.Printf("%b, %d, %o, %x\n", e, e, e)

	fmt.Printf("%.2f\n", math.Pi)
}
