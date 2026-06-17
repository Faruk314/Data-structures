package greedy

import "sort"

func minOperations(nums []int) int {
	ops := 0

	for i := 0; i <= len(nums)-3; i++ {
		if nums[i] == 0 {
			ops++

			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}

	for _, v := range nums {
		if v == 0 {
			return -1
		}
	}

	return ops
}

func FindBuildings(heights []int) []int {
	result := []int{}
	maxHeight := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append(result, i)
			maxHeight = heights[i]
		}
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return result
}

func minimumLength(s string) int {
	frequencies := [26]int{}

	for i := 0; i < len(s); i++ {
		frequencies[s[i]-'a']++
	}

	total := 0

	for _, count := range frequencies {
		if count == 0 {
			continue
		}

		if count%2 == 0 {
			total += 2
		} else {
			total += 1
		}

	}

	return total
}

func MinimumSteps(s string) int64 {
	var zeroCount int64 = 0
	var total int64 = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroCount++
		} else {
			total += zeroCount
		}
	}

	return total
}

func MinIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			diff := nums[i-1] - nums[i] + 1
			moves += diff
			nums[i] = nums[i] + diff
		}
	}

	return moves
}

func maxAbsoluteSum(nums []int) int {
	maxSum := 0
	minSum := 0
	prefix := 0

	for i := 0; i < len(nums); i++ {
		prefix += nums[i]

		if prefix > maxSum {
			maxSum = prefix
		}

		if prefix < minSum {
			minSum = prefix
		}
	}

	return maxSum - minSum
}
