package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 2)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Wassup")
	time.Sleep(time.Millisecond * 10)
}
