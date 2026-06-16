package greedy

import (
	"cmp"
	"slices"
	"sort"
)

func BuyChoco(prices []int, money int) int {
	sort.Ints(prices)
	minCost := prices[0] + prices[1]

	if money >= minCost {
		return money - minCost
	}

	return money
}

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			fives++

		case 10:
			if fives == 0 {
				return false
			}
			fives--
			tens++

		case 20:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	moves := 0

	for idx, seat := range seats {
		moves += abs(seat - students[idx])
	}

	return moves
}

func maximumOddBinaryNumber(s string) string {
	runes := []rune(s)

	slices.SortFunc(runes, func(a, b rune) int {
		return cmp.Compare(b, a)
	})

	lastOneIdx := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == '1' {
			lastOneIdx = i
		}
	}

	runes[len(runes)-1], runes[lastOneIdx] = runes[lastOneIdx], runes[len(runes)-1]

	return string(runes)
}

func maxDepth(s string) int {
	count := 0
	maxCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++

			if count > maxCount {
				maxCount = count
			}
		} else if s[i] == ')' {
			count--
		}
	}

	return maxCount
}

func AreAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	var swaps []int

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			swaps = append(swaps, i)

			if len(swaps) > 2 {
				return false
			}
		}
	}

	return len(swaps) == 2 && s1[swaps[0]] == s2[swaps[1]] && s1[swaps[1]] == s2[swaps[0]]
}
