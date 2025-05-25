package main

import "fmt"

func main() {
	ids := []int{33, 65, 13, 89, 42, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", (i + 1), id)
	}

	// Not using index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", id, i)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s", k)
	}
}
