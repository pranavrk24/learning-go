package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Let's have a look at the type conversions")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value = %v Type = %T", num, num)
}
