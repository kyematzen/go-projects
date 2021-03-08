package main

import (
	"fmt"
)

func main() {
	array := []string{".kye", ".noah", ".dan", ".alex", ".jarret", ".mark"}

	for index, value := range array {
		fmt.Println("index:", index, "value:", value)
	}
}