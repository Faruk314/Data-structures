package stack

import (
	"strconv"
)

func CalPoints(operations []string) int {
	stack := Stack{}
	totalSum := 0

	for _, char := range operations {
		if num, err := strconv.Atoi(char); err == nil {
			stack.Push(num)
			totalSum += num
		} else if char == "+" && len(stack) >= 2 {
			first, _ := stack.Pop()
			second, _ := stack.Pop()

			score := first + second

			stack.Push(second)
			stack.Push(first)
			stack.Push(score)

			totalSum += score

		} else if char == "C" {
			num, _ := stack.Pop()
			totalSum -= num
		} else if char == "D" {
			num, _ := stack.Pop()

			double := num * 2

			stack.Push(num)
			stack.Push(double)

			totalSum += double
		}
	}

	return totalSum
}
