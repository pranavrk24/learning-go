package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func main() {
	a, b := 3.12, 1.784
	fmt.Println(add(a, b))
}

