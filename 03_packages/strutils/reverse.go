package strutils

func ReverseString(s string) string {
	// Convert string to a slice of runes to handle multi-byte characters safely
	runes := []rune(s)

	// Use two pointers starting from opposite ends and swap elements
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	return string(runes)
}
