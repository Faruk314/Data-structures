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

func MergeAlternately(word1, word2 string) string {
	var result strings.Builder

	n1, n2 := len(word1), len(word2)

	maxLen := n1

	if n2 > n1 {
		maxLen = n2
	}

	for i := 0; i < maxLen; i++ {

		if i < n1 {
			result.WriteByte(word1[i])
		}

		if i < n2 {
			result.WriteByte(word2[i])
		}
	}

	return result.String()
}
