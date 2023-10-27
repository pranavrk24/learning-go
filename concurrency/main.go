package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	a := []int{2, 3, 5, 8, 4}

	go printCount(ch)

	for _, v := range a {
		ch <- v
	}
	time.Sleep(time.Millisecond * 2)
	fmt.Println("End of main")
}
