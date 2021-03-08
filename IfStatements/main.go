package main

import (
	"fmt"
)

func main() {

	x := 11

	if x < 2 {
		fmt.Println("2 >", x)
	} else if x > 9 {
		fmt.Println("9 <", x)
	} else {
		fmt.Println("failed to be less than 2 or greater than 9 for value", x)
	}
}