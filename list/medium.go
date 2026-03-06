package list

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
