package main

import "fmt"

func main() {
	// Arrays

	// var fruitArray [2]string

	// // Assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// fruitArray = [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)

	fruitSlice := [3]string{"Apple", "Orange", "Grape"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
