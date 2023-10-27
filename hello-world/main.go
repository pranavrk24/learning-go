package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (0.5 * acceleration * t * t) + (initialVelocity * t) + initialDisplacement
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter acceleration, initial velocity, initial displacement: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	values := strings.Split(text, " ")

	fmt.Println("Enter time: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	time, _ := strconv.ParseFloat(text, 64)

	a, b, c := values[0], values[1], values[2]
	acceleration, _ := strconv.ParseFloat(a, 64)
	initialVelocity, _ := strconv.ParseFloat(b, 64)
	initialDisplacement, _ := strconv.ParseFloat(c, 64)

	displacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(displacement(time))
}
