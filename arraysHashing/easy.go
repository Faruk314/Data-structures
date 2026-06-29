package arrayshashing

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

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

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	var builder strings.Builder

	limit := len(first)

	if len(last) < limit {
		limit = len(last)
	}

	for i := 0; i < limit; i++ {
		if first[i] != last[i] {
			break
		}

		builder.WriteByte(first[i])
	}

	return builder.String()
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 1

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left += 1
		}
	}

	return left
}

func GetConcatenation(nums []int) []int {
	n := len(nums)

	ans := make([]int, 2*n)

	copy(ans[0:n], nums)
	copy(ans[n:], nums)

	return ans
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	j := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}

	return false
}

func ScoreOfString(s string) int {
	score := 0

	for i := 1; i < len(s); i++ {
		score += abs(int(s[i]) - int(s[i-1]))
	}

	return score
}

func CountSeniors(details []string) int {
	count := 0

	for i := 0; i < len(details); i++ {
		detail := details[i]

		ageStr := detail[11:13]

		ageNum, err := strconv.Atoi(ageStr)
		if err != nil {
			continue
		}

		if ageNum > 60 {
			count++
		}
	}

	return count
}

func lengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == ' ' {

			if count > 0 {
				return count
			}

			continue
		}

		count++
	}

	return count
}

func StringMatching(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}

func maxProductDifference(nums []int) int {
	max1, max2 := math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	return (max1 * max2) - (min1 * min2)
}

func GenerateRows(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)

		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}

	return triangle
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		space := flowerbed[i]

		if space == 1 {
			continue
		}

		leftEmpty := i == 0 || flowerbed[i-1] == 0
		rightEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if leftEmpty && rightEmpty {
			flowerbed[i] = 1
			n--

			if n == 0 {
				return true
			}
		}
	}

	return false
}

func MaxDifference(s string) int {
	freq := [26]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxOdd := 0
	minEven := len(s) + 1

	for _, count := range freq {
		if count == 0 {
			continue
		}

		if count%2 != 0 {
			if count > maxOdd {
				maxOdd = count
			}
		} else {
			if count < minEven {
				minEven = count
			}
		}
	}

	return maxOdd - minEven
}

func maxNumberOfBalloons(text string) int {
	var counts [26]int

	for i := 0; i < len(text); i++ {
		counts[text[i]-'a']++
	}

	b := counts['b'-'a']
	a := counts['a'-'a']
	l := counts['l'-'a'] / 2
	o := counts['o'-'a'] / 2
	n := counts['n'-'a']

	return min(b, a, l, o, n)
}

func min(vars ...int) int {
	minVal := vars[0]

	for _, count := range vars {
		if count < minVal {
			minVal = count
		}
	}

	return minVal
}

func MaxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for idx, num := range nums {
		totalSum -= num

		if totalSum == leftSum {
			return idx
		}

		leftSum += num
	}

	return -1
}

func kthDistinct(arr []string, k int) string {
	distinct := make(map[string]int)

	for _, char := range arr {
		distinct[char]++
	}

	count := 0

	for _, str := range arr {
		if distinct[str] == 1 {
			count++

			if count == k {
				return str
			}
		}
	}

	return ""
}

func findDisappearedNumbers(nums []int) []int {
	result := []int{}

	for _, num := range nums {

		targetIndex := abs(num) - 1

		if nums[targetIndex] > 0 {
			nums[targetIndex] = -nums[targetIndex]
		}
	}

	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func LongestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 1
	count := 1
	isIncreasing := true

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {

			if !isIncreasing {
				isIncreasing = true
				count = 1
			}

			count++
		} else if nums[i] < nums[i-1] {

			if isIncreasing {
				isIncreasing = false
				count = 1
			}

			count++
		} else {
			count = 1
		}

		if count > longest {
			longest = count
		}
	}

	return longest
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	totalNumbers := n * n

	freq := make([]int, totalNumbers+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			freq[grid[i][j]]++
		}
	}

	res := make([]int, 2)
	foundRepeated := false
	foundMissing := false

	for i := 1; i <= totalNumbers; i++ {
		if freq[i] == 2 {
			res[0] = i
			foundRepeated = true
		} else if freq[i] == 0 {
			res[1] = i
			foundMissing = true
		}

		if foundRepeated && foundMissing {
			break
		}
	}

	return res
}

func validWordSquare(words []string) bool {
	for r := 0; r < len(words); r++ {
		for c := 0; c < len(words[r]); c++ {

			if c >= len(words) || r >= len(words[c]) {
				return false
			}

			if words[r][c] != words[c][r] {
				return false
			}
		}
	}
	return true
}

func replaceElements(arr []int) []int {
	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > max {
			arr[i], max = max, arr[i]
		} else {
			arr[i] = max
		}
	}

	return arr
}
