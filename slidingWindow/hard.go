package slidingwindow

func minWindow(s string, t string) string {
	mapS := make([]int, 256)
	mapT := make([]int, 256)

	for i := 0; i < len(t); i++ {
		mapT[t[i]]++
	}

	result := ""
	right := 0
	minLen := len(s) + 1

	for left := 0; left < len(s); left++ {

		for right < len(s) && !isDesirable(mapS, mapT) {
			mapS[s[right]]++
			right++
		}

		if isDesirable(mapS, mapT) && right-left < minLen {
			result = s[left:right]
			minLen = right - left
		}

		mapS[s[left]]--
	}

	return result
}

func isDesirable(mapS []int, mapT []int) bool {
	for i := 0; i < 256; i++ {
		if mapT[i] > mapS[i] {
			return false
		}
	}
	return true
}

func help(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	freq := make(map[int]int)
	count := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		for len(freq) > k {
			freq[nums[left]]--

			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}

			left++
		}

		count += right - left + 1
	}

	return count
}

func subarraysWithKDistinct(nums []int, k int) int {
	return help(nums, k) - help(nums, k-1)
}
