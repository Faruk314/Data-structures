package greedy

import (
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
