package main

import "fmt"

func returnValues(x int) {
	x *= 3
	fmt.Println(x)
}

func mutate(sli []int) {
	for i := range sli {
		sli[i] *= 2
	}
	fmt.Println(sli)
}

func main() {
	// x := 2
	// returnValues(x)
	// fmt.Println(x)

	a := []int{1, 2, 3, 4, 5, 6}
	mutate(a)
	fmt.Println(a)
}
