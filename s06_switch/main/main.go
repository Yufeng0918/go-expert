package main

import "fmt"

func main() {
	grade := 'A'

	switch grade {
	case 'A', 'D':
		fmt.Println("Excellent")
	case 'B':
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	score := 100
	switch {
	case score > 50:
		fmt.Println("50")
		fallthrough
	case score >= 80:
		fmt.Println("80")
	}
}
