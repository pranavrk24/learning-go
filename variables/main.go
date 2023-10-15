package main

import "fmt"

func returnValues() (x, y int) {
	x = 32
	x += 32
	y = 64
	y += 32
	return
}

func main() {
	var isAdmin string
	fmt.Println(isAdmin)

	fmt.Println(returnValues())
}
