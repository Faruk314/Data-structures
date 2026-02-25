package main

import (
	"fmt"
	"golang/list"
)

func main() {
    fmt.Println("Hello")

    myList := list.LinkedList{} 

    myList.Prepend(10)
    myList.Prepend(20)
    myList.Append(30)


    myList.Display()
}