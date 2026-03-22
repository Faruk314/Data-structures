package twopointers

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func ValidPalindrome(s string) bool {
	cleaned := nonAlphanumericRegex.ReplaceAllString(s, "")

	word := strings.ToLower(cleaned)

	fmt.Println(word)

	left := 0

	right := len(word) - 1

	for left < right {
		if word[left] != word[right] {
			return false
		}

		left++
		right--
	}

	return true
}
