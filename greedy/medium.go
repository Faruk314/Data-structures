package greedy

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
