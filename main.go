package main

import (
	"fmt"
	arrayshashing "golang/arraysHashing"
)

func main() {
	anagrams := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	result := arrayshashing.GroupAnagrams(anagrams)

	fmt.Println(result, "res")
}
