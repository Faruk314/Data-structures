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