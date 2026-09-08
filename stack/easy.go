package stack

import (
	"strconv"
	"strings"
	"unicode"
)

func CalPoints(operations []string) int {
	stack := Stack[int]{}
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

func ValidParentheses(s string) bool {
	stack := Stack[rune]{}
	pairs := map[string]struct{}{
		"()": {},
		"[]": {},
		"{}": {},
	}

	for _, bracket := range s {
		if bracket == '(' || bracket == '[' || bracket == '{' {
			stack.Push(bracket)
		} else {

			prevBracket, ok := stack.Pop()

			if !ok {
				return false
			}

			var res strings.Builder
			res.WriteRune(prevBracket)
			res.WriteRune(bracket)

			if _, exists := pairs[res.String()]; !exists {
				return false
			}

		}
	}

	return stack.IsEmpty()
}

func minOperations(logs []string) int {
	depth := 0

	for _, log := range logs {
		if log == "../" {
			if depth > 0 {
				depth--
			}
		} else if log == "./" {
			continue
		} else {
			depth++
		}
	}

	return depth
}

func finalPrices(prices []int) []int {
	stack := Stack[int]{}

	for idx, price := range prices {

		for !stack.IsEmpty() {
			top, _ := stack.Peek()

			if prices[top] < price {
				break
			}

			prices[top] = prices[top] - price
			stack.Pop()
		}

		stack.Push(idx)
	}

	return prices
}

func makeGood(s string) string {
	st := Stack[byte]{}

	for i := 0; i < len(s); i++ {

		if !st.IsEmpty() {
			top, _ := st.Peek()

			if top != s[i] && unicode.ToLower(rune(top)) == unicode.ToLower(rune(s[i])) {
				st.Pop()
				continue
			}
		}

		st.Push(s[i])
	}

	var result strings.Builder

	for i := 0; i < len(st); i++ {
		result.WriteByte(st[i])
	}

	return result.String()
}

func minLength(s string) int {
	st := Stack[byte]{}

	for i := 0; i < len(s); i++ {

		if !st.IsEmpty() {
			top, _ := st.Peek()

			if top == 'A' && s[i] == 'B' || top == 'C' && s[i] == 'D' {
				st.Pop()
				continue
			}
		}

		st.Push(s[i])
	}

	return len(st)
}

func clearDigits(s string) string {
	st := Stack[byte]{}

	for i := 0; i < len(s); i++ {

		if !st.IsEmpty() {
			top, _ := st.Peek()

			if top >= 97 && s[i] < 97 {
				st.Pop()
				continue
			}
		}

		st.Push(s[i])
	}

	var result strings.Builder

	for i := 0; i < len(st); i++ {
		result.WriteByte(st[i])
	}

	return result.String()
}
