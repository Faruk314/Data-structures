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
