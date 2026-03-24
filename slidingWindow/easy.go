package slidingwindow

func MaxSizeSubarray(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
