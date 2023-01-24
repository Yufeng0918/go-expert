package main

import "fmt"

func main() {

	var a1 [2]int
	fmt.Printf("%v\n", a1)

	var a2 [2]int = [2]int{2, 3}
	fmt.Printf("%v\n", a2)
	fmt.Printf("%T\n", a2)

	var a3 = [5]bool{0: true, 4: true}
	fmt.Printf("%v\n", a3)

	var a4 = [5]int{1, 4, 9, 16, 25}
	fmt.Printf("%v\n", a4)
	a4[0] = 2
	a4[4] = 50
	fmt.Printf("%v\n", a4)

	for k, v := range a4 {
		fmt.Printf("%v:%v\n", k, v)
	}

	for i := 0; i < len(a4); i++ {
		fmt.Printf("%v-%v\n", i, a4[i])
	}
}
