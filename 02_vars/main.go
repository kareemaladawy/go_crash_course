package main

import "fmt"

// You can declare a var outside a function (gloabl)
var hasACat = true

// You cannot declare a shorthand variable outside a function
// hasCar := false -> error

func main() {
	// Using var to declare variables
	var name = "Kareem"
	var age byte = 25 // byte: alias for uint8 (unsigned int 8)
	const isCool = true
	var size float32 = 1.5

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

	// Using shorthand to declare variables
	color, email := "Green", "email.com"

	fmt.Println(name, age, color, hasACat, size, email)
}
