package main

import (
	"fmt"
	"go-expert/user"
)

func main() {
	s := user.Hello()
	fmt.Println("Hello, World, " + s)
}
