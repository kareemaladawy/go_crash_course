package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func getSum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greet("James"))
	fmt.Println(getSum(12, 3))
}
