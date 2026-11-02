package main

import "fmt"

func main() {
	// Range with array/slice
	// A slice
	ids := []int{2, 4, 8, 9}

	// _ = unused but has to get declare
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Sum ids
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	// Range with map
	emails := map[int]string{
		1: "bob@e.com",
		2: "sara@e.com",
		3: "mike@e.com",
	}

	for k, v := range emails {
		fmt.Printf("%d: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("ID: %d\n", k)
	}
}
