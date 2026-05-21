package twopointers

import (
	"regexp"
	"sort"
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

func MoveZeroes(nums []int) []int {
	last_non_zero_idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[last_non_zero_idx] = nums[i]
			last_non_zero_idx++
		}
	}

	for i := last_non_zero_idx; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func TwoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	maxSum := -1

	for left < right {
		newSum := nums[left] + nums[right]

		if newSum < k {
			left++

			if newSum > maxSum {
				maxSum = newSum
			}
		} else {
			right--
		}

	}

	return maxSum
}

func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 0
	right := n - 1
	p := n - 1

	for left <= right {

		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[p] = leftSquare
			left++
		} else {
			result[p] = rightSquare
			right--
		}

		p--
	}

	return result
}
