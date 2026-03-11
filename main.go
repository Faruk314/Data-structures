package main

import (
	"fmt"
	arrayshashing "golang/arraysHashing"
)

func main() {
	numbers := []int{}

	result := arrayshashing.TwoSum(numbers, 40)

	fmt.Println(result, "res")
}
