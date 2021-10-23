package main

import (
	"fmt"
)

// float64 is the type if it's not specified
func add(a, b float64) float64 {
	return a + b
}

func main() {
	a, b := 3.12, 1.784
	fmt.Println(add(a, b))
}

