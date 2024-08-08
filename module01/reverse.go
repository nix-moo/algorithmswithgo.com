package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	drow := ""

	// loop from front
	for i := 0; i <= len(word); i++ {
		drow = string(word[i]) + drow
	}
	// loop from back
	for i := len(word) - 1; i >= 0; i-- {
		drow = drow + string(word[i])
	}

	// use runes instead of bytes
	for _, r := range word {
		drow = string(r) + drow
	}

	return drow
}

// more efficient - Builder
func EfficientReverse(word string) string {
	var built strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		built.WriteByte(word[i])
	}
	return built.String()
}
