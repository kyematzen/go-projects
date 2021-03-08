package main

import (
	"fmt"
)

func main() {
	// x holds 5 integer values
	// starts at 0 index

	var x [5]int

	// input value of 11 to position 3
	x[2] = 11

	// y holds 5 integer values
	y := [5]int{5,4,3,2,1}

	fmt.Println(x)
	fmt.Println(y)
}