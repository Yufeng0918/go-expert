package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func print() {
	r := strings.NewReader("hello world")
	r1 := strings.NewReader("hello second")
	io.Copy(os.Stdout, r)

	buf := make([]byte, 8)
	io.CopyBuffer(os.Stdout, r1, buf)
	fmt.Printf("\n%v\n", string(buf))
}

func main() {
	print()
}
