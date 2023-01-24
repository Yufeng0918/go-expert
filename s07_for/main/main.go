package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}

	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("%v %v\n", i, v)
	}

	m := make(map[string]string, 0)
	m["name"] = "jack"
	m["email"] = "abc@gmail.com"
	for k, v := range m {
		fmt.Printf("%v %v\n", k, v)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%v", i)
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		if i == 8 {
			break
		}
		fmt.Printf("%v", i)
	}
}
