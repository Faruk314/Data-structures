package slidingwindow

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
