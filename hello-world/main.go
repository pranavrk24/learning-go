package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello world")
	var slc = []int{}
	slc = append(slc, 2, 33, 123, 8, 3)
	fmt.Println(slc)
	sort.Ints(slc)
	fmt.Println(slc)
}
