package main

import (
	"fmt"
)

type Animals struct {
	name string
}

func (c Animals) eat() {
	fmt.Printf("%v eat\n", c.name)
}

func (c Animals) sleep() {
	fmt.Printf("%v sleep\n", c.name)
}

type Dog struct {
	Animals
	color string
}

func newDog(name string, color string) (*Dog, error) {
	if name == "" {
		return nil, fmt.Errorf("name is error")
	}

	var dog = Dog{Animals{
		name: name,
	},
		color,
	}

	return &dog, nil
}

func main() {

	var c0 = Dog{
		Animals{
			name: "wangcai",
		},
		"blue",
	}

	c0.eat()
	c0.sleep()

	var d0, _ = newDog("huahua", "red")
	fmt.Printf("%v\n", *d0)
}
