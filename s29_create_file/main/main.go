package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file name: %v\n", (*f).Name())
	}
}

func makeDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("tmp: %v\n", s)
}

func remove() {
	os.Remove("a.txt")
	fmt.Printf("remove a.txt\n")
	os.RemoveAll("a/b/c")
	fmt.Printf("remove a/b/c\n")
}

func readFile() {
	b, _ := os.ReadFile("dao/userDao.go")
	fmt.Printf("b: %v\n", string(b[:]))
}

func writeFile() {
	os.WriteFile("test2.txt", []byte("hello"), os.ModePerm)
}

func main() {

	createFile()
	makeDir()
	remove()
	wd()
	readFile()
	writeFile()
}
