package arrayshashing

func HasDuplicates(numbers []int) bool {
	numberSet := make(map[int]struct{})

	for _, number := range numbers {
		if _, ok := numberSet[number]; ok {
			return true
		} else {
			numberSet[number] = struct{}{}
		}
	}

	return false
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, r := range s {
		counts[r]++
	}

	for _, r := range t {
		counts[r]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {

		complement := target - num

		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}

		seen[num] = i

	}

	return nil
}
