package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
)

func tcp() {
	conn, err := net.Dial("tcp", "golang.org:80")
	panicErr(err)

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	panicErr(err)
	fmt.Println(status)
}

func httpGET() {
	resp, err := http.Get("https://golang.org")
	panicErr(err)

	body, err := io.ReadAll(resp.Body)
	panicErr(err)

	fmt.Println(string(body))
}

func main() {
	httpGET()
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}
