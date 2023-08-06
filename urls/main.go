package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Hello")
	path := &url.URL{
		Host:       "8080",
		OmitHost:   false,
		ForceQuery: true,
	}

	fmt.Println(path)
}
