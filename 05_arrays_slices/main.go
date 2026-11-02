package main

import "fmt"

func main() {
	// Arrays in Go have a fixed length
	// Declare
	// fruitArr := [2]string

	// Assign
	// fruitArr[0] = "Banana"
	// fruitArr[1] = "Banana"

	// Declare and assign in the same line
	// fruitArr := [2]string{"Apple", "Orange"}

	// Slices in Go are arrays but without a fixed length
	fruitSlice := []string{"Orange", "Banana", "Apple"}

	fmt.Println(len(fruitSlice)) // 3
	fmt.Println(fruitSlice[0:2]) // returns items in indexes 0 & 1
	fmt.Println(fruitSlice[1])
	fmt.Println(fruitSlice[2])
}
