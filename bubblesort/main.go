package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(sli []int) {
	size := len(sli)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, j int) {
	sli[j], sli[j+1] = sli[j+1], sli[j]
}

func main() {
	fmt.Println("Enter a sequence of 10 integers separated by space: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fields := strings.Fields(text)

	arr := make([]int, len(fields))
	for i, j := range fields {
		arr[i], _ = strconv.Atoi(j)
	}

	BubbleSort(arr)
	fmt.Println(arr)
}
