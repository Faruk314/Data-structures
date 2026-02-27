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