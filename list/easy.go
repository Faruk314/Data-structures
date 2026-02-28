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