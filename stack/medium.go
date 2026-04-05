package stack

import (
	"sort"
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

func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	helperStack := Stack[int]{}

	for idx := n - 1; idx >= 0; idx-- {

		for {
			topIdx, exists := helperStack.Peek()
			if !exists || temperatures[idx] < temperatures[topIdx] {
				break
			}
			helperStack.Pop()
		}

		if topIdx, exists := helperStack.Peek(); exists {
			result[idx] = topIdx - idx
		}

		helperStack.Push(idx)
	}

	return result
}

type car struct {
	pos  int
	time float64
}

func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	cars := make([]car, n)
	for i := 0; i < n; i++ {

		t := float64(target-position[i]) / float64(speed[i])
		cars[i] = car{pos: position[i], time: t}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	stack := Stack[float64]{}

	for _, c := range cars {

		if stack.IsEmpty() {
			stack.Push(c.time)
			continue
		}

		leadTime, _ := stack.Peek()

		if c.time > leadTime {
			stack.Push(c.time)
		}

	}

	return len(stack)
}
