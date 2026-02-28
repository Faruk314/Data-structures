package main

import (
	"fmt"
	"golang/list"
)
func main() {
    list1 := list.LinkedList{}

   
    list1.Append(4)
    list1.Append(1)
    list1.Append(3)
    list1.Append(4)
    list1.Append(5)

    fmt.Print("Original: ")
    
    list1.Display()

    result := list.RemoveElements(list1.Head, 4)

    fmt.Println("result", result.Value)
}