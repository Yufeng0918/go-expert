package main

import "fmt"

func main() {

	m1 := map[string]string{"name": "tom", "age": "20"}
	fmt.Printf("%v\n", m1)

	for k, v := range m1 {
		fmt.Printf("%v-> %v \n", k, v)
	}

	m2 := make(map[string]string)
	m2["name"] = "jack"
	m2["age"] = "20 "
	v, k := m2["name"]
	fmt.Printf("%v-> %v \n", k, v)
	fmt.Printf("%v\n", m2)
}
