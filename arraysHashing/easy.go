package arrayshashing

import (
	"sort"
	"strings"
)

func HasDuplicates(numbers []int) bool {
	numberSet := make(map[int]struct{})

	for _, number := range numbers {
		if _, ok := numberSet[number]; ok {
			return true
		} else {
			numberSet[number] = struct{}{}
		}
	}

	return false
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, r := range s {
		counts[r]++
	}

	for _, r := range t {
		counts[r]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {

		complement := target - num

		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}

		seen[num] = i

	}

	return nil
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	var builder strings.Builder

	limit := len(first)

	if len(last) < limit {
		limit = len(last)
	}

	for i := 0; i < limit; i++ {
		if first[i] != last[i] {
			break
		}

		builder.WriteByte(first[i])
	}

	return builder.String()
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 1

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left += 1
		}
	}

	return left
}
