package main

import "fmt"

func main() {
	a := 100

	if a > 200 {
		fmt.Println("> 200")
	} else if a > 100 {
		fmt.Println("> 100")
	} else {
		fmt.Println("<= 100")
	}

}
