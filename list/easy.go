package list

func MergeList(l1 *LinkedList, l2 *LinkedList) *Node {
	if l1 == nil || l1.Head == nil {
		return l2.Head
	}
	if l2 == nil || l2.Head == nil {
		return l1.Head
	}

	dummy := &Node{}
	tail := dummy
	p1, p2 := l1.Head, l2.Head

	for p1 != nil && p2 != nil {
		if p1.Value <= p2.Value {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}

	return dummy.Next
}

func Reverselist(head *Node) *Node {
	current := head

	var prev *Node

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}

	return prev
}

func RemoveElements(head *Node, value int) *Node {
	dummy := &Node{Value: 0, Next: head}

	current := dummy

	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}

func isPalindrome(list *LinkedList) bool {
	if list.Head == nil || list.Head.Next == nil {
		return true
	}

	head := list.Head
	fast := head
	slow := head

	// 1. Find the middle (slow will be at the start of second half)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 2. Reverse the second half of the list
	var prev *Node // In Go, use the specific pointer type or nil
	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	// 3. Compare the two halves
	left := head
	right := prev // 'prev' is now the head of the reversed second half

	for right != nil {
		if left.Value != right.Value {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func RemoveDuplicateFromSortedList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Value == curr.Next.Value {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}

	return head
}

func GetIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := 0, 0
	currA, currB := headA, headB

	for currA != nil {
		lenA++
		currA = currA.Next
	}
	for currB != nil {
		lenB++
		currB = currB.Next
	}

	currA, currB = headA, headB

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			currA = currA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			currB = currB.Next
		}
	}

	for currA != currB {
		currA = currA.Next
		currB = currB.Next
	}

	return currA
}
