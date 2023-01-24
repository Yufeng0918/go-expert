package main

import "fmt"

type Person struct {
	id, age     int
	name, email string
}

func main() {

	var p1 Person
	p1.id = 1
	p1.name = "p1"
	fmt.Printf("%v\n", p1)

	var p2 = Person{
		id:   2,
		age:  20,
		name: "jack",
	}
	fmt.Printf("%v\n", p2)

	var p3 = Person{
		2,
		21,
		"mike",
		"mike@gmail.com",
	}
	fmt.Printf("%v\n", p3)

	ap3 := &p3
	fmt.Printf("%T\n", ap3)
	(*ap3).id = 3
	fmt.Printf("%v\n", p3)
}
