package main

import (
    "fmt"
)

func main() {
    i := 7 // start at 7
    inc(&i) // attempts to increase by 1
    fmt.Println(i) // prints 8
}

func inc(x *int) { // grabs variable's memory location.
    *x++ // increase's variable value from local memory location, not from method's variable.
    // usually without returning it should output 7.
}