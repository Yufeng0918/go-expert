package main

import "fmt"

type Dog struct {
	name  string
	color string
}

type Person struct {
	id, age     int
	name, email string
}

type Male struct {
	id   int
	name string
	dog  Dog
}

type pPerson *Person

func showPerson(p Person) {
	p.id = 0
	p.name = "jack"
	fmt.Printf("%v\n", p)
}

func changePerson(p *Person) {

	(*p).id = 0
	(*p).name = "jack"
	fmt.Printf("%v\n", *p)
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

	var p_person *Person = &p3
	fmt.Printf("%T\n", p_person)
	fmt.Printf("%v\n", *p_person)

	showPerson(p3)
	fmt.Printf("%v\n", p3)

	changePerson(&p3)
	fmt.Printf("%v\n", p3)

	var m1 = Male{
		2,
		"mike",
		Dog{
			"wangcai",
			"white",
		},
	}
	fmt.Printf("%v\n", m1)

}
