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
