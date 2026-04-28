package arrayshashing

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		chars := strings.Split(s, "")

		sort.Strings(chars)

		key := strings.Join(chars, "")

		groups[key] = append(groups[key], s)

	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func TopKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := []int{}

	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	result := make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	visited := make(map[int]bool)
	for _, num := range nums {
		visited[num] = false
	}

	longest := 0

	for _, num := range nums {

		if visited[num] {
			continue
		}

		visited[num] = true
		currentLength := 1

		for next := num + 1; ; next++ {
			if _, exists := visited[next]; exists {
				visited[next] = true
				currentLength++
			} else {
				break
			}
		}

		for prev := num - 1; ; prev-- {
			if _, exists := visited[prev]; exists {
				visited[prev] = true
				currentLength++
			} else {
				break
			}
		}

		if currentLength > longest {
			longest = currentLength
		}
	}

	return longest
}

func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := SortArray(nums[:mid])
	right := SortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func IsValidSudoku(board [][]byte) bool {
	if checkHorisontal(board) == false {
		return false
	}
	if checkVertical(board) == false {
		return false
	}

	subBox := 0
	startRow := 0
	startCol := 0

	for subBox < 9 {
		subBoxSet := make(map[byte]struct{})

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				num := board[startRow+r][startCol+c]
				if num == '.' {
					continue
				}
				if _, exists := subBoxSet[num]; exists {
					return false
				}
				subBoxSet[num] = struct{}{}
			}
		}

		startCol += 3
		if startCol == 9 {
			startCol = 0
			startRow += 3
		}
		subBox++
	}
	return true
}

func checkHorisontal(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		horisontalSet := make(map[byte]struct{})
		for col := 0; col < 9; col++ {
			num := board[row][col]

			if num == '.' {
				continue
			}

			if _, exists := horisontalSet[num]; exists {
				return false
			}

			horisontalSet[num] = struct{}{}
		}
	}

	return true
}

func checkVertical(board [][]byte) bool {
	for col := 0; col < 9; col++ {
		verticalSet := make(map[byte]struct{})
		for row := 0; row < 9; row++ {

			num := board[row][col]

			if num == '.' {
				continue
			}

			if _, exists := verticalSet[num]; exists {
				return false
			}

			verticalSet[num] = struct{}{}

		}
	}

	return true
}

func vowelStrings(words []string, queries [][]int) []int {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	prefix := make([]int, len(words)+1)

	for i, word := range words {
		if vowels[word[0]] && vowels[word[len(word)-1]] {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		left, right := q[0], q[1]
		ans[i] = prefix[right+1] - prefix[left]
	}

	return ans
}
