package main

import "fmt"

func main() {
	x := 5
	y := &x // memory address of x
	fmt.Println(x, &x, y, *y)

	*y = 10
	fmt.Println(x)

	*y = *y * *y

	fmt.Println(x)
	fmt.Println(*y)
}
