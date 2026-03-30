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
