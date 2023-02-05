package main

import (
	"io"
	"os"
)

func OpenClose() {
	f, _ := os.Open("test2 .txt")
	//fmt.Printf("file name: %v\n", f.Name())
	f.Close()
}

func readOps() {
	f, _ := os.Open("dao/userDao.go")
	buf := make([]byte, 10)
	for {
		_, err := f.Read(buf)

		if err == io.EOF {
			break
		}
		fmt.Printf("%v", string(buf))
	}
	fmt.Println()

	de, _ := os.ReadDir("dao")
	for _, v := range de {
		fmt.Printf("name: %v\n", v.Name())
	}
}

func main() {
	OpenClose()
	readOps()
}
