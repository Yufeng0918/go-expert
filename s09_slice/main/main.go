package main

import "fmt"

func main() {

	var a1 = make([]int, 2)
	fmt.Printf("%v\n", a1)

	var a2 = []int{2, 3}
	fmt.Printf("%v\n", a2)
	fmt.Printf("%T\n", a2)

	var a3 = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", a3[0:3])
	fmt.Printf("%v\n", a3[3:])
	fmt.Printf("%v\n", a3[:])

	var a4 = a3[4:]
	fmt.Printf("%v\n", a4)

	for i, v := range a4 {
		fmt.Printf("%v->%v\t", i, v)
	}
	fmt.Printf("\n")

	var a5 []int
	a5 = append(a5, 100)
	fmt.Printf("%v\n", a5[:])
	a5 = append(a5, a3[2:]...)
	fmt.Printf("%v\n", a5[:])
	a5 = append(a3[:2], a3[3:]...)
	fmt.Printf("%v\n", a5[:])

	var a6 = make([]int, 5)
	copy(a6, a5)
	a6[0] = 100
	fmt.Printf("%v\n", a5[:])
	fmt.Printf("%v\n", a6[:])
}
