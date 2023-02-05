package main

import (
	"os"
)

func OpenClose() {
	//f, _ := os.OpenFile("test2.txt", os.O_RDWR, 0755)
	f, _ := os.OpenFile("test2.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.Write([]byte("hello golang\n"))
	f.Close()
}

func main() {
	OpenClose()
}
