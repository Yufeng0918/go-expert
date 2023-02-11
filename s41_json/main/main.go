package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id, Age     int
	Name, Email string
}

func main() {

	p := Person{Id: 1, Age: 20, Name: "jack", Email: "jack@gmail.com"}
	b, _ := json.Marshal(p)
	fmt.Printf("person %v\n", string(b))

	b0 := []byte(`{"Id":2,"Age":20,"Name":"jack","Email":"jack@gmail.com"}`)
	var p0 Person
	json.Unmarshal(b0, &p0)
	fmt.Printf("person %v\n", p0)
}
