package main

import "fmt"

func main() {
	// Maps are key -> value pairs
	// Declare
	// emails := make(map[int]string)

	// Assign
	// emails[1] = "bob@e.com"
	// emails[2] = "sara@e.com"
	// emails[3] = "mike@e.com"

	// Declare & assign
	emails := map[int]string{
		1: "bob@e.com",
		2: "sara@e.com",
		3: "mike@e.com",
	}

	fmt.Println(len(emails))
	fmt.Println(emails)

	delete(emails, 2)
	fmt.Println(emails)
}
