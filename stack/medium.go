package stack

import (
	"strconv"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(v int) {
	s.stack = append(s.stack, v)

	if len(s.minStack) == 0 || v <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, v)
	}
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}

	topVal := s.stack[len(s.stack)-1]
	if topVal == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}

	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

func EvalRPN(tokens []string) int {
	myStack := Stack[int]{}

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			myStack.Push(num)
		} else {
			a, _ := myStack.Pop()
			b, _ := myStack.Pop()

			switch token {
			case "+":
				myStack.Push(b + a)
			case "-":
				myStack.Push(b - a)

			case "*":
				myStack.Push(b * a)
			case "/":
				myStack.Push(b / a)
			}
		}
	}

	result, _ := myStack.Pop()
	return result
}

func AsteroidCollision(asteroids []int) []int {
	Stack := Stack[int]{}

	for _, asteroid := range asteroids {
		isAlive := true

		for isAlive && !Stack.IsEmpty() && asteroid < 0 {
			last, _ := Stack.Peek()

			if last < 0 {
				break
			}

			if -asteroid > last {
				Stack.Pop()
			} else if -asteroid == last {

				Stack.Pop()
				isAlive = false
			} else {
				isAlive = false
			}
		}

		if isAlive {
			Stack.Push(asteroid)
		}
	}

	return Stack
}
