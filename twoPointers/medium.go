package twopointers

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func MaxArea(heights []int) int {
	leftPointer := 0
	rightPointer := len(heights) - 1
	biggestSurface := 0

	for leftPointer < rightPointer {

		leftHeight := heights[leftPointer]

		rightHeight := heights[rightPointer]

		height := min(leftHeight, rightHeight)

		width := rightPointer - leftPointer

		surface := height * width

		if surface > biggestSurface {
			biggestSurface = surface
		}

		if rightHeight >= leftHeight {
			leftPointer++
		} else {
			rightPointer--
		}

	}

	return biggestSurface
}

func BoatsToSavePeople(people []int, limit int) int {
	slices.Sort(people)
	count := 0
	left := 0
	right := len(people) - 1

	for left <= right {

		if people[left]+people[right] <= limit {
			left++
		}

		right--
		count++
	}

	return count
}

func LongestMountain(arr []int) int {
	res := 0

	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			left, right := i, i

			for left > 0 && arr[left] > arr[left-1] {
				left--
			}

			for right+1 < len(arr) && arr[right] > arr[right+1] {
				right++
			}

			total := right - left + 1

			if total > res {
				res = total
			}
		}
	}

	return res
}

func DividePlayers(skills []int) int64 {
	n := len(skills)
	sort.Ints(skills)

	target := skills[0] + skills[n-1]
	var totalChemistry int64 = 0

	left := 0
	right := n - 1

	for left < right {
		sum := skills[left] + skills[right]

		if sum != target {
			return -1
		}

		totalChemistry += int64(skills[left]) * int64(skills[right])

		left++
		right--
	}

	return totalChemistry
}

func MinCost(colors string, neededTime []int) int {
	time := 0
	left := 0
	right := 1

	for right < len(colors) {
		if colors[right] != colors[left] {
			left = right
			right++
		} else {
			if neededTime[left] < neededTime[right] {
				time += neededTime[left]
				left = right
			} else {
				time += neededTime[right]
			}

			right++
		}
	}

	return time
}

func RearrangeArray(nums []int) []int {
	positive := []int{}
	negative := []int{}

	for _, n := range nums {
		if n > 0 {
			positive = append(positive, n)
		} else {
			negative = append(negative, n)
		}
	}

	i := 0

	for 2*i < len(nums) {

		nums[2*i] = positive[i]
		nums[2*i+1] = negative[i]

		i += 1
	}

	return nums
}

func BagOfTokens(tokens []int, power int) int {
	sort.Ints(tokens)
	score := 0
	maxScore := 0
	left := 0
	right := len(tokens) - 1

	for left <= right {
		if power >= tokens[left] {
			score++
			power -= tokens[left]
			left++

			if score > maxScore {
				maxScore = score
			}

		} else if score > 0 {
			power += tokens[right]
			score--
			right--
		} else {
			break
		}
	}

	return score
}

func RearrangeArrayNotEqualToAverage(nums []int) []int {
	sort.Ints(nums)
	res := []int{}

	left, right := 0, len(nums)-1

	for len(res) != len(nums) {

		res = append(res, nums[left])
		left++

		if left <= right {
			res = append(res, nums[right])
			right--
		}

	}

	return res
}

func pivotArray(nums []int, pivot int) []int {
	var less, equal, greater []int

	for _, n := range nums {
		switch {
		case n < pivot:
			less = append(less, n)
		case n == pivot:
			equal = append(equal, n)
		default:
			greater = append(greater, n)
		}
	}

	result := append(less, equal...)
	result = append(result, greater...)

	return result
}

func MaxWidthRamp(nums []int) int {
	maxRight := make([]int, len(nums))
	prevMax := 0

	for i := len(nums) - 1; i >= 0; i-- {
		maxRight[i] = max(nums[i], prevMax)
		prevMax = maxRight[i]

	}

	res := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		for nums[left] > maxRight[right] {
			left++
		}

		res = max(res, right-left)
	}

	return res
}

func AddSpaces(s string, spaces []int) string {
	res := make([]byte, len(s)+len(spaces))
	sIdx := 0
	spaceIdx := 0

	for i := 0; i < len(res); i++ {
		if spaceIdx < len(spaces) && sIdx == spaces[spaceIdx] {
			res[i] = ' '
			spaceIdx++
		} else {
			res[i] = s[sIdx]
			sIdx++
		}
	}

	return string(res)
}

func numSubseq(nums []int, target int) int {
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	mod := 1000000007

	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}

	sort.Ints(nums)

	left := 0
	right := n - 1
	count := 0

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			count = (count + pow2[right-left]) % mod
			left++
		}
	}

	return count
}

func MinimumLength(s string) int {
	left := 0
	right := len(s) - 1

	for left < right && s[left] == s[right] {
		ch := s[left]

		for left <= right && s[left] == ch {
			left++
		}

		for left <= right && s[right] == ch {
			right--
		}
	}

	if left > right {
		return 0
	}

	return right - left + 1
}

func KthGrammar(n int, k int) int {
	left := 1
	right := 1 << (n - 1)

	current := 0

	for left < right {
		mid := left + (right-left)/2

		if k <= mid {
			right = mid
		} else {
			if current == 0 {
				current = 1
			} else {
				current = 0
			}

			left = mid + 1
		}
	}

	return current
}

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	left := 0
	charsIdx := 0

	for left < len(chars) {
		right := left

		for right < len(chars) && chars[left] == chars[right] {
			right++
		}

		total := right - left

		chars[charsIdx] = chars[left]
		charsIdx++

		if total > 1 {
			str := strconv.Itoa(total)

			for i := 0; i < len(str); i++ {
				chars[charsIdx] = str[i]
				charsIdx++
			}
		}

		left = right
	}

	return charsIdx
}

func ProductOfTwoEncoded(encoded1, encoded2 [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0

	for i < len(encoded1) && j < len(encoded2) {
		val1, freq1 := encoded1[i][0], encoded1[i][1]
		val2, freq2 := encoded2[j][0], encoded2[j][1]

		prod := val1 * val2
		minFreq := min(freq1, freq2)

		if len(result) > 0 && result[len(result)-1][0] == prod {
			result[len(result)-1][1] += minFreq
		} else {
			result = append(result, []int{prod, minFreq})
		}

		encoded1[i][1] -= minFreq

		if encoded1[i][1] == 0 {
			i++
		}

		encoded2[j][1] -= minFreq

		if encoded2[j][1] == 0 {
			j++
		}
	}

	return result
}

func MeetingScheduler(slots1, slots2 [][]int, duration int) []int {
	i, j := 0, 0

	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})

	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	for i < len(slots1) && j < len(slots2) {
		startTime := max(slots1[i][0], slots2[j][0])
		endTime := min(slots1[i][1], slots2[j][1])

		if endTime-duration >= startTime {
			return []int{startTime, startTime + duration}
		}

		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}

	}

	return []int{}
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}

	n1, n2 := len(words1), len(words2)
	left := 0
	right := 0

	for left < n2 && words1[left] == words2[left] {
		left++
	}

	for right < n2-left && words1[n1-1-right] == words2[n2-1-right] {
		right++
	}

	return left+right == n2
}

func reverseWords(s string) string {
	words := strings.Fields(s)

	left := 0
	right := len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]

		left++
		right--
	}

	return strings.Join(words, " ")
}

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	max := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			i++
		} else {

			if j-i > max {
				max = j - i
			}

			j++
		}
	}

	return max
}

func minPrice(prices []int, discounts []int) float64 {
	sort.Ints(prices)
	sort.Ints(discounts)

	var total float64
	j := len(discounts) - 1

	for i := len(prices) - 1; i >= 0; i-- {
		if j >= 0 {
			total += float64(prices[i]) * (100 - float64(discounts[j])) / 100
			j--
		} else {
			total += float64(prices[i])
		}
	}

	return total
}

func reverse(bytes []byte) string {
	left := 0
	right := len(bytes) - 1

	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}

	return string(bytes)
}

func getVowelCount(word string) int {
	vowelCount := 0
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for i := 0; i < len(word); i++ {
		if vowels[word[i]] {
			vowelCount++
		}
	}

	return vowelCount
}

func reverseWordsWithSameVowels(s string) string {
	words := strings.Fields(s)
	totalCount := getVowelCount(words[0])

	for i := 1; i < len(words); i++ {
		count := getVowelCount(words[i])

		if count == totalCount {
			bytes := []byte(words[i])
			reversed := reverse(bytes)
			words[i] = reversed
		}
	}

	return strings.Join(words, " ")
}

func getBound(nums []int, bound int) int {
	count := 0
	ans := 0

	for _, num := range nums {
		if num <= bound {
			count++
		} else {
			count = 0
		}
		ans += count
	}

	return ans
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return getBound(nums, right) - getBound(nums, left-1)
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		curr := matrix[row][col]

		if curr == target {
			return true
		} else if curr > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	i, j := 0, 0

	sort.Ints(trainers)
	sort.Ints(players)

	matches := 0

	for i < len(players) && j < len(trainers) {
		if players[i] > trainers[j] {
			j++
		} else {

			matches++

			i++
			j++
		}
	}

	return matches
}
