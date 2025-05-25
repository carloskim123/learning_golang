package main

import "fmt"

func main() {
	// Define Map
	// emails := make(map[string]string)

	// // Assign kv
	// emails["Bob"] = "bog@mail.com"
	// emails["Sharon"] = "sharon@mail.com"
	// emails["Mike"] = "mike@mail.com"

	// Delcare map and add Key Values (kv)
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
}
