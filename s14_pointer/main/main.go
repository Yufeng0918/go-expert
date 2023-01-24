package main

import "fmt"

func main() {

	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	*ip = 200
	fmt.Printf("ip: %v\n", i)

	a := [3]int{1, 2, 3}
	var ap [3]*int
	fmt.Printf("len: %v\n", len(a))
	for i := 0; i < len(a); i++ {
		ap[i] = &a[i]
	}
	fmt.Printf("ap: %v\n", ap)
}
