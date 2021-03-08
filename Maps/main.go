package main

import (
	"fmt"
)

func main() {
	examplePriority := make(map[string]int)

	examplePriority["a"] = 0
	examplePriority["b"] = 1
	examplePriority["c"] = 2

	fmt.Println(examplePriority)

	delete(examplePriority, "b")

	fmt.Println(examplePriority)
}