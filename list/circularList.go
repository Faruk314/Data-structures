package list

import (
	"fmt"
)

func (l *LinkedList) PrependCircular(val int) {
    newNode := &Node{Value: val}
    
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.Next = l.Head
		newNode.Prev = l.Tail
        
		l.Tail.Next = newNode
		l.Head.Prev = newNode

		l.Head = newNode

       
    }
    l.Size++
}

func (l *LinkedList) AppendCircular(val int) {
    newNode := &Node{Value: val}

    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
       
        newNode.Next = newNode
        newNode.Prev = newNode
    } else {
        newNode.Next = l.Head
        newNode.Prev = l.Tail

        l.Tail.Next = newNode
        l.Head.Prev = newNode

        l.Tail = newNode
    }
    l.Size++
}

func (l *LinkedList) DeleteHeadCircular() {
    if l.Head == nil {
        return
    }

    if(l.Head == l.Tail) {
        l.Head = nil
        l.Tail = nil
    
    }else {
    newHead := l.Head.Next

     newHead.Prev = l.Tail

     l.Tail.Next = newHead

     l.Head.Next = nil
     l.Head.Prev = nil

     l.Head = newHead  
    }

    l.Size--     
}

func (l *LinkedList) DeleteTailCircular() {
    if l.Head == nil {
        return
    }

    if l.Head == l.Tail {
        l.Head = nil
        l.Tail = nil
    } else {
        newTail := l.Tail.Prev
        newTail.Next = l.Head
        l.Head.Prev = newTail

      
        l.Tail.Next = nil
        l.Tail.Prev = nil
        l.Tail = newTail
    }

    l.Size-- 
}


func (l *LinkedList) InsertAtIndexCircular(val int, index int) {

    if index < 0 || index > l.Size {
        fmt.Println("Index is out of range")

        return
    }

    if index == 0 {
        l.PrependCircular(val)
        return
    }

    if index == l.Size {
        l.AppendCircular(val)
        return
    }

    current := l.Head
    newNode := &Node{Value: val}

    for i:=0; i < index - 1; i++ {
        current = current.Next
    }

    newNode.Next = current.Next
    newNode.Prev = current

    current.Next.Prev = newNode

    current.Next = newNode

    l.Size++
}
