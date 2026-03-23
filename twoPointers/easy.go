package twopointers

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func ValidPalindrome(s string) bool {
	cleaned := nonAlphanumericRegex.ReplaceAllString(s, "")

	word := strings.ToLower(cleaned)

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

func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func ValidPalindromeTwo(s string) bool {
	cleaned := nonAlphanumericRegex.ReplaceAllString(s, "")
	word := strings.ToLower(cleaned)

	left := 0
	right := len(word) - 1

	for left < right {
		if word[left] == word[right] {
			left++
			right--
		} else {
			return isPalindrome(word, left+1, right) || isPalindrome(word, left, right-1)
		}
	}

	return true
}
