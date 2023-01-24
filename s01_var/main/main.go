package main

import "fmt"

func getNameAndAge() (string, int) {
	return "mick", 20
}

func main() {
	var name string
	var surname = "jack"
	var age int
	var isGood bool = true
	const PI float64 = 3.14

	var myname, myage, b = "tom", 20, true

	name = "abc"
	age = 123
	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(age)
	fmt.Println(isGood)
	fmt.Printf("name %v, age %v, b %v", myname, myage, b)

	_, age1 := getNameAndAge()
	fmt.Println("get age %v", age1)
	fmt.Println(PI)
}
