package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	s := "abc"
	site := WebSite{
		Name: "360",
	}

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", site)
	fmt.Printf("%T\n", site)

	i := 96
	fmt.Printf("%c\n", i)

	x := 100
	p := &x
	fmt.Printf("%p\n", p)
}
