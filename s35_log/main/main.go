package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Printf("my log %d", 100)
	log.Fatalln("my fatal")
}

func test2() {
	log.Panicln("my panic`")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("mY LOG: ")
	f, _ := os.OpenFile("a.log", os.O_RDWR|os.O_CREATE, 0777)
	log.Fatalln("fatal")
	log.SetOutput(f)
}

func main() {

	test1()
	test2()
	fmt.Println("end")
}
