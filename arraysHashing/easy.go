package arrayshashing

import (
	"sort"
	"strconv"
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

func GetConcatenation(nums []int) []int {
	n := len(nums)

	ans := make([]int, 2*n)

	copy(ans[0:n], nums)
	copy(ans[n:], nums)

	return ans
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	j := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}

	return false
}

func ScoreOfString(s string) int {
	score := 0

	for i := 1; i < len(s); i++ {
		score += abs(int(s[i]) - int(s[i-1]))
	}

	return score
}

func CountSeniors(details []string) int {
	count := 0

	for i := 0; i < len(details); i++ {
		detail := details[i]

		ageStr := detail[11:13]

		ageNum, err := strconv.Atoi(ageStr)
		if err != nil {
			continue
		}

		if ageNum > 60 {
			count++
		}
	}

	return count
}

func lengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == ' ' {

			if count > 0 {
				return count
			}

			continue
		}

		count++
	}

	return count
}

func StringMatching(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}
