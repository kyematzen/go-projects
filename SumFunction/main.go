package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2,3))
}

func sum(x int, y int) int {
	return x + y
}