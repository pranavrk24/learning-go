package main

import "fmt"

// this is how you define a struct in go
type dog struct {
	breed  string
	color  string
	weight float32
	price  float32
}

// value receiver
func (d dog) inr() float32 {
	discount := 10
	d.price += 10 // do not modify the struct's actual value
	return d.price - d.price*float32(discount)/100
}

// pointer receiver
func (d *dog) usd() float32 {
	discount := 10
	d.price += 10 // do modify the struct's actual value
	return d.price - d.price*float32(discount)/100
}

func main() {
	/* can be initialised like below
	   pakkun := dog{"pomerian","white",35.8,15000.5}
	*/
	pakkun := dog{
		breed:  "pomerian",
		color:  "white",
		weight: 35.8,
		price:  15000.5}
	fmt.Println(pakkun.breed, pakkun.color, pakkun.price, pakkun.weight)
	fmt.Println(pakkun.inr())
	fmt.Println(pakkun.breed, pakkun.color, pakkun.price, pakkun.weight)
	fmt.Println(pakkun.usd())
	fmt.Println(pakkun.breed, pakkun.color, pakkun.price, pakkun.weight)
}
