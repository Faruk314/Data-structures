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
