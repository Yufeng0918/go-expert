package main

import "fmt"

type Person struct {
	id, age     int
	name, email string
}

func main() {

	var tom Person
	tom.id = 1
	tom.name = "tom"

	fmt.Printf("%v\n", tom)
}
