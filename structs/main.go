package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver -> Greeting method
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
	// Init Person using struct
	person1 := Person{"Carlos", "Kirui", "Baltimore", "m", 16}

	fmt.Println(person1)

	person1.age++
	fmt.Println(person1.age)

	fmt.Println(person1.greet())
}
