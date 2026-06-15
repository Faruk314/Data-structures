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
