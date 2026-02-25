package main

import (
	"fmt"
	"golang/list"
)

func main() {
    fmt.Println("Hello")

    myList := list.LinkedList{} 

    myList.PrependDoubly(20);

    myList.AppendDoubly(40)


    myList.Display()
}