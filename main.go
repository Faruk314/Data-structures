package main

import (
	"fmt"
	arrayshashing "golang/arraysHashing"
)

func main() {
	numbers := []int{10, 20, 30, 40}

	result := arrayshashing.HasDuplicates(numbers)

	fmt.Println(result, "res")
}
