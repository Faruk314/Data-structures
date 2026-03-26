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
