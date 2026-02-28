package list


func RemoveElements(head *Node, value int) *Node {
   
   dummy := &Node{Value: 0, Next: head}

   current := dummy

    for current.Next != nil {

        if current.Next.Value == value {

            current.Next = current.Next.Next
        }else {
            current = current.Next
        }
    }
  

   return dummy.Next
}