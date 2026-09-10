package mathgeometry

import (
	"math"
	"strings"
)

func IntToRoman(num int) string {
	type romanMap struct {
		val int
		sym string
	}

	romanSymbols := []romanMap{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for _, entry := range romanSymbols {
		for num >= entry.val {
			num -= entry.val
			result.WriteString(entry.sym)
		}
	}

	return result.String()
}

func checkPowersOfThree(n int) bool {
	var currNum float64
	nums := []float64{}
	idx := 0

	for {
		currNum = math.Pow(3, float64(idx))

		if currNum > float64(n) {
			break
		}

		nums = append(nums, currNum)
		idx++
	}

	totalSum := nums[len(nums)-1]

	if totalSum == float64(n) {
		return true
	}

	for i := len(nums) - 2; i >= 0; i-- {

		if totalSum+nums[i] > float64(n) {
			continue
		}

		totalSum += nums[i]

		if totalSum == float64(n) {
			return true
		}
	}

	return false
}
