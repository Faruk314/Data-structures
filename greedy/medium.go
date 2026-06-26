package greedy

import (
	"sort"
)

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

func MaxAbsoluteSum(nums []int) int {
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

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		currentMax = max(nums[i], currentMax+nums[i])

		globalMax = max(currentMax, globalMax)
	}

	return globalMax
}

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	oddCount := 0
	for i := 0; i < 26; i++ {
		if freq[i]%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= k
}

func MaxSubarraySumCircular(nums []int) int {
	currMax, globalMax := nums[0], nums[0]
	currMin, globalMin := nums[0], nums[0]

	totalSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], currMax+nums[i])
		globalMax = max(currMax, globalMax)

		currMin = min(nums[i], currMin+nums[i])
		globalMin = min(currMin, globalMin)

		totalSum += nums[i]
	}

	if globalMax < 0 {
		return globalMax
	}

	return max(globalMax, totalSum-globalMin)
}

func minSwaps(nums []int) int {
	totalOnes := 0

	for _, num := range nums {
		if num == 1 {
			totalOnes += 1
		}
	}

	if totalOnes <= 1 {
		return 0
	}

	nums = append(nums, nums...)
	minSwaps := len(nums)
	currentZeroes := 0
	left := 0

	for right := 0; right < len(nums); right++ {

		if nums[right] == 0 {
			currentZeroes++
		}

		if right-left+1 > totalOnes {

			if nums[left] == 0 {
				currentZeroes--
			}

			left++
		}

		if right-left+1 == totalOnes {
			if currentZeroes < minSwaps {
				minSwaps = currentZeroes
			}
		}

	}

	return minSwaps
}

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	left := 0
	longest := 1
	var prev string

	for right := 1; right < len(arr); right++ {
		if arr[right-1] > arr[right] && prev != ">" {
			prev = ">"
			longest = max(longest, right-left+1)

		} else if arr[right-1] < arr[right] && prev != "<" {
			prev = "<"
			longest = max(longest, right-left+1)

		} else if arr[right-1] == arr[right] {

			left = right

			prev = ""
		} else {
			left = right - 1
		}
	}

	return longest
}

func canJump(nums []int) bool {
	topIdx := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= topIdx {
			topIdx = i
		}
	}

	return topIdx == 0
}

func jump(nums []int) int {
	n := len(nums)
	destination := n - 1
	coverage := 0
	lastJumpIdx := 0
	totalJumps := 0

	if n == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		coverage = max(coverage, i+nums[i])

		if i == lastJumpIdx {
			lastJumpIdx = coverage
			totalJumps++

			if coverage >= destination {
				return totalJumps
			}
		}
	}

	return totalJumps
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalCost > totalGas {
		return -1
	}

	startIndex := 0
	currGas := 0

	for i := 0; i < len(gas); i++ {
		currGas += gas[i] - cost[i]

		if currGas < 0 {
			currGas = 0
			startIndex = i + 1
		}
	}

	return startIndex
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	freq := make(map[int]int)

	for _, card := range hand {
		freq[card]++
	}

	uniques := make([]int, 0, len(freq))
	for card := range freq {
		uniques = append(uniques, card)
	}

	sort.Ints(uniques)

	for _, card := range uniques {
		for freq[card] > 0 {
			for next := card; next < card+groupSize; next++ {
				if freq[next] == 0 {
					return false
				}
				freq[next]--
			}
		}
	}

	return true
}

func minChanges(s string) int {
	changes := 0

	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			changes++
		}
	}

	return changes
}
