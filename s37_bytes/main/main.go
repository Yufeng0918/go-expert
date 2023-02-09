package main

import (
	"bytes"
	"fmt"
)

func main() {

	var s string = "openai.com"
	b := []byte(s)
	b1 := []byte(".com")
	fmt.Printf("b => %v\n", string(b))
	fmt.Printf("b contains %v\n", bytes.Contains(b, b1))
	fmt.Printf("b repeat %v\n", string(bytes.Repeat(b1, 3)))

	s0 := []byte("你好世界")
	r := bytes.Runes(s0)
	fmt.Printf("len s0 => %v\n", len(s0))
	fmt.Printf("len r => %v\n", len(r))

}
