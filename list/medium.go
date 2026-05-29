package list

import "golang/stack"

func SumLists(list1, list2 LinkedList) LinkedList {
	result := LinkedList{}

	curr1 := list1.Head
	curr2 := list2.Head
	carry := 0

	for curr1 != nil || curr2 != nil || carry != 0 {
		sum := carry

		if curr1 != nil {
			sum += curr1.Value
			curr1 = curr1.Next
		}

		if curr2 != nil {
			sum += curr2.Value
			curr2 = curr2.Next
		}

		carry = sum / 10
		newValue := sum % 10

		result.Append(newValue)
	}

	return result
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	firstStack := stack.Stack[*Node]{}
	secondStack := stack.Stack[*Node]{}

	curr1 := l1
	for curr1 != nil {
		firstStack.Push(curr1)
		curr1 = curr1.Next
	}

	curr2 := l2
	for curr2 != nil {
		secondStack.Push(curr2)
		curr2 = curr2.Next
	}

	var head *Node
	carry := 0

	for !firstStack.IsEmpty() || !secondStack.IsEmpty() || carry > 0 {
		sum := carry

		if n1, ok := firstStack.Pop(); ok {
			sum += n1.Value
		}

		if n2, ok := secondStack.Pop(); ok {
			sum += n2.Value
		}

		remainder := sum % 10
		carry = sum / 10

		resultNode := &Node{Value: remainder}
		resultNode.Next = head
		head = resultNode
	}

	return head
}

func RemoveNthFromEnd(l *LinkedList, n int) *Node {
	if l.Head == nil {
		return nil
	}

	dummy := &Node{Next: l.Head}

	left := dummy
	right := l.Head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}

func SwapPairs(l *LinkedList) *Node {
	if l.Head == nil {
		return nil
	}

	dummy := &Node{Value: 0, Next: l.Head}

	point := dummy

	for point.Next != nil && point.Next.Next != nil {

		swap1 := point.Next
		swap2 := point.Next.Next

		swap1.Next = swap2.Next
		swap2.Next = swap1

		point.Next = swap2
		point = swap1

	}

	return dummy.Next
}

func PartionList(head *Node, x int) *Node {
	leftHead := &Node{}
	rightHead := &Node{}

	left := leftHead
	right := rightHead

	curr := head

	for curr != nil {
		if curr.Value < x {
			left.Next = curr
			left = left.Next
		} else {
			right.Next = curr
			right = right.Next

		}
		curr = curr.Next
	}

	right.Next = nil

	left.Next = rightHead.Next

	return leftHead.Next
}

func CopyListWithRandomPointer(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldToCopy := make(map[*Node]*Node)

	curr := head

	for curr != nil {
		oldToCopy[curr] = &Node{Value: curr.Value}
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		copy := oldToCopy[curr]

		copy.Next = oldToCopy[curr.Next]
		copy.Random = oldToCopy[curr.Random]

		curr = curr.Next
	}

	return oldToCopy[head]
}

func RotateList(head *Node, k int) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	n := k % length
	if n == 0 {
		return head
	}

	tail.Next = head

	curr := head
	for i := 0; i < length-n-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil

	return newHead
}

func LinkedListCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func SwapNodes(head *Node, k int) *Node {
	if head == nil || k <= 0 {
		return head
	}

	fast := head

	for i := 1; i < k; i++ {

		if fast.Next == nil {
			return head
		}

		fast = fast.Next
	}

	first := fast

	slow := head

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	last := slow

	first.Value, last.Value = last.Value, first.Value

	return head
}

func modifiedList(nums []int, head *Node) *Node {
	numSet := make(map[int]struct{})

	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	dummy := &Node{Next: head}

	prev := dummy
	curr := head

	for curr != nil {
		if _, ok := numSet[curr.Value]; ok {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

func swapPairs(head *Node) *Node {
	dummy := &Node{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		first.Next = second.Next
		second.Next = first
		prev.Next = second

		prev = first
	}

	return dummy.Next
}

func removeNodes(head *Node) *Node {
	stack := stack.Stack[*Node]{}

	curr := head

	for curr != nil {

		for {
			top, ok := stack.Peek()

			if !ok || top.Value >= curr.Value {
				break
			}

			stack.Pop()
		}

		stack.Push(curr)
		curr = curr.Next
	}

	var resultHead *Node

	for !stack.IsEmpty() {
		node, _ := stack.Pop()

		node.Next = resultHead
		resultHead = node

	}

	return resultHead
}

func pairSum(head *Node) int {
	slow, fast := head, head
	var slowPrev *Node

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slowPrev = slow
		slow = slow.Next
	}

	if slowPrev != nil {
		slowPrev.Next = nil
	}

	var prev *Node
	curr := slow

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	p1 := head
	p2 := prev
	max := 0

	for p2 != nil {

		sum := p1.Value + p2.Value

		if sum > max {
			max = sum
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return max
}

func mergeInBetween(list1 *Node, a int, b int, list2 *Node) *Node {
	curr := list1

	for i := 1; i < a; i++ {
		curr = curr.Next
	}

	start := curr
	lastPrev := start

	for i := 0; i <= (b - a); i++ {
		lastPrev = lastPrev.Next
	}

	end := lastPrev.Next
	lastPrev.Next = nil

	start.Next = list2

	curr = list2

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = end

	return list1
}
