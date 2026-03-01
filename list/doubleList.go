package list

import "fmt"

func (l *LinkedList) PrependDoubly(val int) {
	newNode := &Node{Value: val}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head

		l.Head.Prev = newNode

		l.Head = newNode
	}
	l.Size++
}

func (l *LinkedList) AppendDoubly(val int) {
	newNode := &Node{Value: val}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {

		newNode.Prev = l.Tail

		l.Tail.Next = newNode

		l.Tail = newNode
	}
	l.Size++
}

func (l *LinkedList) DeleteHeadDoubly() {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
	l.Size--

	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Prev = nil
	}
}

func (l *LinkedList) DeleteTailDoubly() {
	if l.Tail == nil {
		return
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {

		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	}

	l.Size--
}

func (l *LinkedList) DisplayDoubly() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}

	fmt.Println("nil")
}

func (l *LinkedList) InsertAtIndexDoubly(val int, index int) {
	if index < 0 || index > l.Size {
		fmt.Println("Index is out of range")
		return
	}

	if index == 0 {
		l.PrependDoubly(val)
		return
	}

	if index == l.Size {
		l.AppendDoubly(val)
		return
	}

	newNode := &Node{Value: val}
	current := l.Head

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Prev = current
	newNode.Next = current.Next

	current.Next.Prev = newNode
	current.Next = newNode

	l.Size++
}
