package main

import (
	"fmt"
	"sort"
)

func main() {

	f := []float64{1.1, 4.4, 5.5, 2.2}
	sort.Float64s(f)
	fmt.Println("%v\n", f)

	ls := sort.StringSlice{"100", "42", "41"}
	fmt.Println("%v\n", ls)

}
