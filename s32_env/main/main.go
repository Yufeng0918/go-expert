package main

import (
	"fmt"
	"os"
)

func showEnv() {
	s, _ := os.LookupEnv("JAVA_HOME")
	fmt.Printf("java home: %v\n", s)

	os.Setenv("foo", "bar")
	e, _ := os.LookupEnv("foo")
	fmt.Printf("foo: %v\n", e)
}

func main() {
	showEnv()
}
