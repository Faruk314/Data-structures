package main

import (
	"fmt"
	twopointers "golang/twoPointers"
)

func main() {
	mySlice := []int{1, 2, 1}
	result := twopointers.Palindrome(mySlice)
	fmt.Println("result", result)
}
