package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for x := 1; x <= 10; x++ {
		fmt.Printf("Number %d\n", x)
	}

	// Fizz Buzz
	for j := 1; j <= 100; j++ {
		fizz := j%3 == 0
		buzz := j%5 == 0

		if fizz && buzz {
			fmt.Printf("%d: Fizz Buzz\n", j)
		} else if fizz {
			fmt.Printf("%d: Fizz\n", j)
		} else if buzz {
			fmt.Printf("%d: Buzz\n", j)
		}
	}
}
