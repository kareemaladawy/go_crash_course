package main

import (
	"fmt"
	"math"

	"github.com/kareemaladawy/go_crash_course/03_packages/strutils" // A package we created
)

func main() {
	fmt.Println(math.Floor(2.78))
	fmt.Println(math.Ceil(2.78))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutils.ReverseString("olleH"))
}
