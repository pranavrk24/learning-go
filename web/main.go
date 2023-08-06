package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	resp, _ := http.Get("https://lco.dev")
	fmt.Println(resp)

	sum := 0
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}

	fmt.Println(sum)
	// close the response
	defer resp.Body.Close()
}
