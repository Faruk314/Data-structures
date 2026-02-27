package main

import (
	"fmt"
	"golang/list"
)

func main() {
    fmt.Println("Hello")

    list1 := list.LinkedList{} 

    list1.Append(9)
    list1.Append(2)

    list2 := list.LinkedList{}
    list2.Append(5)
    list2.Append(5)

    result := list.SumLists(list1, list2)

    result.Display()
    
}