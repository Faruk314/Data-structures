package slidingwindow

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	left := 0
	charSet := make(map[byte]struct{})
	maxLen := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		for {
			_, exists := charSet[char]

			if !exists {
				break
			}

			delete(charSet, s[left])
			left++
		}

		charSet[s[right]] = struct{}{}

		currLength := right - left + 1

		if currLength > maxLen {
			maxLen = currLength
		}

	}

	return maxLen
}

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	left := 0
	maxFreq := 0
	maxWindow := 0

	for right := 0; right < len(s); right++ {

		charIdx := s[right] - 'A'
		freq[charIdx]++

		if freq[charIdx] > maxFreq {
			maxFreq = freq[charIdx]
		}

		windowLength := right - left + 1

		if windowLength-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}

		currentWindow := right - left + 1
		if currentWindow > maxWindow {
			maxWindow = currentWindow
		}
	}

	return maxWindow
}

func CheckInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	var s1Count, s2Count [26]int

	for i := 0; i < n1; i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	if s1Count == s2Count {
		return true
	}

	for i := n1; i < n2; i++ {

		s2Count[s2[i]-'a']++
		s2Count[s2[i-n1]-'a']--

		if s1Count == s2Count {
			return true
		}
	}

	return false
}

func MinSubArrayLen(nums []int, target int) int {
	minWindowSize := math.MaxInt
	currentSum := 0
	right := 0
	left := 0

	for right < len(nums) {
		currentSum += nums[right]
		right++

		for currentSum >= target {
			currentWindowSize := right - left

			if currentWindowSize < minWindowSize {
				minWindowSize = currentWindowSize
			}

			currentSum -= nums[left]
			left++
		}
	}

	return minWindowSize
}

func LongestOnes(nums []int, k int) int {
	maxLen := 0
	num_zeroes := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			num_zeroes += 1
		}

		for num_zeroes > k {
			if nums[left] == 0 {
				num_zeroes -= 1
			}

			left++
		}

		newLen := right - left + 1

		if newLen > maxLen {
			maxLen = newLen
		}
	}

	return maxLen
}

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	base := 0

	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			base += customers[i]
		}
	}

	left := 0
	extra := 0
	maxExtra := 0

	for right := 0; right < len(customers); right++ {
		if grumpy[right] == 1 {
			extra += customers[right]
		}

		if right-left+1 > minutes {
			if grumpy[left] == 1 {
				extra -= customers[left]
			}

			left++
		}

		if extra > maxExtra {
			maxExtra = extra
		}
	}

	return base + maxExtra
}

func TotalFruit(fruits []int) int {
	basket := make(map[int]int)
	maxFruits := 0
	left := 0

	for right := 0; right < len(fruits); right++ {

		basket[fruits[right]]++

		for len(basket) > 2 {
			basket[fruits[left]]--

			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}
			left++
		}

		currentWindowSize := right - left + 1
		if currentWindowSize > maxFruits {
			maxFruits = currentWindowSize
		}
	}

	return maxFruits
}

func MinOperations(nums []int, x int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	target := totalSum - x

	if target == 0 {
		return len(nums)
	}

	if target < 0 {
		return -1
	}

	maxLen := -1
	currentSum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]

		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}

		if currentSum == target {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}
	}

	if maxLen == -1 {
		return -1
	}

	return len(nums) - maxLen
}

func MaxConsTwo(nums []int) int {
	left := 0
	num_zeroes := 0
	max_window := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			num_zeroes += 1
		}

		for num_zeroes > 1 {

			if nums[left] == 0 {
				num_zeroes -= 1
			}

			left++
		}

		curr_window := right - left + 1

		if curr_window > max_window {
			max_window = curr_window
		}
	}

	return max_window
}
