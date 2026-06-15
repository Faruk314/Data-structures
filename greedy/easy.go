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
