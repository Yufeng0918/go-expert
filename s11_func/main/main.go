package main

import "fmt"

func f1() {
	fmt.Printf("hello world\n")
}

func sum(a int, b int) (ret int) {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func compare(a int, b int) bool {
	return a > b
}

func f2() (string, string) {
	return "jack", "20"
}

func f3(s []int) {
	s[0] = 100
}

func f4(args ...int) {
	for _, arg := range args {
		fmt.Printf("%v\n", arg)
	}
}

func sayHello(name string) {
	fmt.Printf("hello, %s\n", name)
}

func sayBye(name string) {
	fmt.Printf("bye, %s\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {

	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return add
	}
}

func main() {

	fmt.Printf("%v\n", sum(1, 2))
	f1()
	fmt.Printf("%v\n", compare(1, 2))
	k, v := f2()
	fmt.Printf("%v, %v\n", k, v)

	s := []int{1, 2, 3}
	f3(s)
	fmt.Printf("%v\n", s)
	f4(1, 2, 3, 4)

	type f func(int, int) int
	var f1 f = sum
	fmt.Printf("%v\n", f1(1, 2))
	f1 = max
	fmt.Printf("%v\n", f1(1, 2))
	test("jack", sayHello)
	test("jack", sayBye)

	f2 := cal("+")
	fmt.Printf("cal %v\n", f2(1, 2))
	f2 = cal("-")
	fmt.Printf("cal %v\n", f2(1, 2))

	s1 := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Printf("any func %v\n", s1)

}
