package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func writeToBuf() {

	s1 := strings.NewReader("hello world")
	br := bufio.NewReader(s1)
	b := bytes.NewBuffer(make([]byte, 0))
	br.WriteTo(b)
	fmt.Printf("write to buffer %c\n", b)
}

func writeToFile() {

	f, _ := os.OpenFile("test.txt", os.O_RDWR, 0777)
	w := bufio.NewWriter(f)
	w.WriteString("hello buffer !!!!!")
}

func writeTo() {

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("1234")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

func scanner() {
	s := strings.NewReader("abc def ghi")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func main() {
	//r := strings.NewReader("HELLO WORLD")
	r, _ := os.Open("dao/userDao.go")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s  %v\n", s)

	s1 := strings.NewReader("hello world")
	br := bufio.NewReader(s1)
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	s2 := strings.NewReader("hello go\nhello java")
	br2 := bufio.NewReader(s2)
	line, _, _ := br2.ReadLine()
	fmt.Printf("%c\n", line)

	writeToBuf()
	writeToFile()
	writeTo()
	scanner()
}
