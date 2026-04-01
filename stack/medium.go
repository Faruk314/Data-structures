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