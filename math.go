package main

import (
	"fmt"
	"math"
)

func test() {
	fmt.Println("This is inside a function")
}

func main() {
	fmt.Println("The square root of 64 is", math.Sqrt(64))
	test()
}
