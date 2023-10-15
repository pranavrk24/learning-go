package main

import "fmt"

func returnValues(x int) {
	x *= 3
	fmt.Println(x)
}

func main() {
	x := 2
	returnValues(x)
	fmt.Println(x)
}
