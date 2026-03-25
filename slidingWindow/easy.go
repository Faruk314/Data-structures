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

func isVowel(letter byte) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func MaxVowels(s string, k int) int {
	if len(s) < k {
		return 0
	}

	currentCount := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currentCount++
		}
	}

	maxCount := currentCount

	if maxCount == k {
		return k
	}

	for i := k; i < len(s); i++ {

		if isVowel(s[i]) {
			currentCount += 1
		}

		if isVowel(s[i-k]) {
			currentCount -= 1
		}

		if currentCount > maxCount {
			maxCount = currentCount

			if maxCount == k {
				return maxCount
			}
		}
	}

	return maxCount
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numSet := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := numSet[nums[i]]; ok {
			return true
		}

		numSet[nums[i]] = struct{}{}

		if len(numSet) > k {
			delete(numSet, nums[i-k])
		}

	}

	return false
}

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buyPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			currentProfit := prices[i] - buyPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}
