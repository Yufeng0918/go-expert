package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("%v computer read\n", c.name)
}

func (c Computer) write() {
	fmt.Printf("%v computer ssd\n", c.name)
}

type Mobile struct {
	name string
}

func (m Mobile) read() {
	fmt.Printf("%v mobile read\n", m.name)
}

func (m Mobile) write() {
	fmt.Printf("%v mobile 5g\n", m.name)
}

func main() {

	var c0 = Computer{
		name: "AMD",
	}
	c0.read()
	c0.write()

	var m0 = Mobile{"iPhone"}
	m0.read()
	m0.write()
}
