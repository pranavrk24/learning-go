package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct {
}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct {
}

func (b Snake) Eat()   { fmt.Println("mice") }
func (b Snake) Move()  { fmt.Println("slither") }
func (b Snake) Speak() { fmt.Println("hsss") }

func createAnimal(animals map[string]Animal, name, action string) {
	var animal Animal
	switch action {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	}

	animals[name] = animal
	fmt.Println("Created it!")
}

func queryAnimal(animals map[string]Animal, name, action string) {
	animal, found := animals[name]
	if !found {
		fmt.Println("Not found!")
		return
	}

	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	animals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		var cmd, name, action string
		fmt.Scanf("%s %s %s", &cmd, &name, &action)
		if strings.ToLower(cmd) == "newanimal" {
			createAnimal(animals, name, action)
		} else if strings.ToLower(cmd) == "query" {
			queryAnimal(animals, name, action)
		} else {
			fmt.Println("Invalid command!")
		}
	}
}
