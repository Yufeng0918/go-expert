package main

import "fmt"

type Person struct {
	id, age     int
	name, email string
}

type Customer struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("%v, eat....\n", per.name)
}

func (per *Person) sleep() {
	(*per).name = "haaa"
	fmt.Printf("%v, sleep....\n", per.name)
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("%v login\n", name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {

	var p2 = Person{
		id:   2,
		age:  20,
		name: "jack",
	}
	p2.eat()
	(&p2).sleep()

	var c1 = Customer{
		"tom",
	}

	fmt.Printf("result %v\n", c1.login("tom", "123"))
	fmt.Printf("result %v\n", c1.login("jack", "123"))
}
