package greedy

import (
	"fmt"
	"golang/stack"
	"math"
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

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)
	n := len(nums)

	right := n - 1
	left := n - 3
	minDiff := math.MaxInt

	for i := 0; i < 4; i++ {
		min := (right + 1) % n
		max := left - 1

		if left-1 < 0 {
			max = n - 1
		}

		if nums[max]-nums[min] < minDiff {
			minDiff = nums[max] - nums[min]
		}

		left = (left + 1) % n
		right = (right + 1) % n
	}

	return minDiff
}

func MaximumImportance(n int, roads [][]int) int64 {
	freq := make([]int, n)

	for _, road := range roads {
		freq[road[0]]++
		freq[road[1]]++
	}

	sort.Ints(freq)

	fmt.Println(freq)

	var totalSum int64

	for i := 1; i <= n; i++ {
		totalSum += int64(freq[i-1]) * int64(i)
	}

	return totalSum
}

func MinimumPushes(word string) int {
	freq := make([]int, 26)

	for _, c := range word {
		freq[c-'a']++
	}

	sort.Ints(freq)

	counter := 0
	multiplier := 1
	total := 0

	for i := 25; i >= 0; i-- {
		counter++
		total += multiplier * freq[i]

		if counter == 8 {
			multiplier += 1
			counter = 0
		}
	}

	return total
}

func PredictPartyVictory(senate string) string {
	radiants := 0
	dires := 0

	for i := 0; i < len(senate); i++ {
		if senate[i] == 'D' {
			dires++
		} else {
			radiants++
		}
	}

	bSenate := []byte(senate)
	banR := 0
	banD := 0
	i := 0

	for radiants > 0 && dires > 0 {
		if bSenate[i] == 'R' {
			if banR > 0 {

				radiants--
				banR--
				bSenate[i] = 'X'
			} else {
				banD++
			}
		} else if bSenate[i] == 'D' {
			if banD > 0 {

				dires--
				banD--
				bSenate[i] = 'X'
			} else {
				banR++
			}
		}

		i = (i + 1) % len(bSenate)
	}

	if radiants > 0 {
		return "Radiant"
	}
	return "Dire"
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	prefixSum := make([]int, n+1)
	suffixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + cardPoints[i]
	}

	for j := 0; j < n; j++ {
		suffixSum[j+1] = suffixSum[j] + cardPoints[n-1-j]
	}

	maxPoints := 0

	for x := 0; x <= k; x++ {
		y := k - x
		points := prefixSum[x] + suffixSum[y]
		if points > maxPoints {
			maxPoints = points
		}
	}

	return maxPoints
}

func mergeTriplets(triplets [][]int, target []int) bool {
	valid := make(map[int]struct{})

	for i := 0; i < len(triplets); i++ {

		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2] {
			continue
		}

		for j := 0; j < 3; j++ {
			if triplets[i][j] == target[j] {
				valid[j] = struct{}{}
			}
		}

	}

	return len(valid) == 3
}

func partitionLabels(s string) []int {
	lastIndex := [26]int{}
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]-'a'] = i
	}

	left := 0
	maxEnd := 0
	result := []int{}

	for right := 0; right < len(s); right++ {
		charLast := lastIndex[s[right]-'a']
		if charLast > maxEnd {
			maxEnd = charLast
		}

		if right == maxEnd {
			result = append(result, right-left+1)
			left = right + 1
		}
	}

	return result
}

func removePairs(s string, target string, score int) (string, int) {
	st := stack.Stack[byte]{}
	gainedScore := 0

	firstChar := target[0]
	secondChar := target[1]

	for i := 0; i < len(s); i++ {
		char := s[i]
		top, exists := st.Peek()

		if exists && top == firstChar && char == secondChar {
			st.Pop()
			gainedScore += score
		} else {
			st.Push(char)
		}
	}

	remaining := make([]byte, len(st))

	for i := len(st) - 1; i >= 0; i-- {
		char, _ := st.Pop()
		remaining[i] = char
	}

	return string(remaining), gainedScore
}

func maxGain(s string, x int, y int) int {
	firstPair, firstScore := "ab", x
	secondPair, secondScore := "ba", y

	if y > x {
		firstPair, firstScore = "ba", y
		secondPair, secondScore = "ab", x
	}

	totalScore := 0

	sAfterFirstPass, score1 := removePairs(s, firstPair, firstScore)
	totalScore += score1

	_, score2 := removePairs(sAfterFirstPass, secondPair, secondScore)
	totalScore += score2

	return totalScore
}

func winnerOfGame(colors string) bool {
	alicePairs := 0
	bobPairs := 0

	left := 0

	for right := 0; right < len(colors); right++ {
		if right-left+1 == 3 {
			pattern := colors[left : right+1]

			if pattern == "AAA" {
				alicePairs++
			} else if pattern == "BBB" {
				bobPairs++
			}

			left++
		}
	}

	if alicePairs > bobPairs {
		return true
	}

	return false
}

func minimumDeletions(s string) int {
	st := stack.Stack[byte]{}
	count := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == 'a' && !st.IsEmpty() {
			st.Pop()
			count++
		} else if char == 'b' {
			st.Push(char)
		}
	}

	return count
}

func minDeletions(s string) int {
	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	usedFreq := make(map[int]struct{})
	res := 0

	for _, count := range freq {
		for count > 0 {
			if _, exists := usedFreq[count]; !exists {
				break
			}

			count--
			res++
		}

		usedFreq[count] = struct{}{}
	}

	return res
}
