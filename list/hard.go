package list

func MergeKLists(lists []*LinkedList) *Node {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var nextRound []*LinkedList

		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]

			if i+1 < len(lists) {
				l2 := lists[i+1]

				mergedHead := MergeList(l1, l2)
				nextRound = append(nextRound, &LinkedList{Head: mergedHead})
			} else {
				nextRound = append(nextRound, l1)
			}
		}

		lists = nextRound
	}

	return lists[0].Head
}

func ReverseKGroup(head *Node, k int) *Node {
	if head == nil || k == 1 {
		return head
	}

	dummy := &Node{Next: head}
	groupPrev := dummy

	for {
		kth := getKth(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		prev := groupNext
		curr := groupPrev.Next

		for i := 0; i < k; i++ {
			nextTmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = nextTmp
		}

		newTail := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = newTail
	}

	return dummy.Next
}

func getKth(curr *Node, k int) *Node {
	for curr != nil && k > 1 {
		curr = curr.Next
		k--
	}
	return curr
}
