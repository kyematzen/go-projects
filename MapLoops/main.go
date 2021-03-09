package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for index, value := range m {
		fmt.Println("index:", index, "value:", value)
	}
}