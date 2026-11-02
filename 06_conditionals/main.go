package main

import "fmt"

func main() {
	x := 89
	y := 10

	// if else
	if x < y { // common practice is to not use parentheses
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// if else if
	color := "Red"

	if color == "Red" {
		fmt.Println("Color is red")
	} else if color == "Blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red nor blue")
	}

	// Switch
	switch color {
	case "Red":
		{
			fmt.Println("Red is the color")
		}
	case "Blue":
		{
			fmt.Println("Blue is the color")
		}
	default:
		{
			fmt.Println("Color is unknown")
		}
	}
}
