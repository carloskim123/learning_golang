package main

import (
	"fmt"
)

var name = " Carlos Kirui"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8  in16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var

	// var name = "Carlos Kirui"
	var age int = 37
	var isCool = true
	isCool = false
	// var size float32 = 2.3

	// Shorthand
	// name := "Carlton"
	// email := "carloskirui154@gmail.com"

	name, email := "Carlos", "carloskirui154@gmail.com"

	fmt.Println(name, age, isCool, email)
	// fmt.Printf("%T\n", )
}
