package list

import "fmt"

func (l *LinkedList) Prepend(val int) {
    newNode := &Node{Value: val}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.Next = l.Head
        l.Head = newNode
    }
    l.Size++
}

func (l *LinkedList) Append(val int) {
    newNode := &Node{Value: val}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode    } else {
        l.Tail.Next = newNode
        l.Tail = newNode
    }
    l.Size++
}

func (l *LinkedList) DeleteHead() {
    if l.Head == nil {
        return
    }
    l.Head = l.Head.Next
    l.Size--
    if l.Head == nil {
        l.Tail = nil 
    }
}

func (l *LinkedList) DeleteTail() {
    if l.Head == nil {
        return
    }
    if l.Head == l.Tail {
        l.Head = nil
        l.Tail = nil
    } else {
        current := l.Head
        for current.Next != l.Tail {
            current = current.Next
        }
        current.Next = nil
        l.Tail = current
    }
    l.Size--
}

func (l *LinkedList) Display() {
    current := l.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

func (l *LinkedList) InsertAtIndex(val int, index int) {
    if index < 0 || index > l.Size {
        fmt.Println("Index out of bounds")
        return
    }

   
    if index == 0 {
        l.Prepend(val)
        return
    }

   
    if index == l.Size {
        l.Append(val)
        return
    }

   
    newNode := &Node{Value: val}
    current := l.Head
    
    
    for i := 0; i < index-1; i++ {
        current = current.Next
    }

    
    newNode.Next = current.Next
  
    current.Next = newNode
    
    l.Size++
}










