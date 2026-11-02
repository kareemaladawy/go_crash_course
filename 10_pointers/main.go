package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)
	// Use * to read the value from the address
	fmt.Println(*b)

	// Changing the value of the address changes the original value
	*b = 9
	fmt.Println(*b, a)

	// Why use pointers? It can increase performance if you pass the address instead of a copy of data
	// or you want to change values at a specific location
}
