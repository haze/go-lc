package main

import "fmt"

func anagramMappings(a, b []int) []int {
	reverseIndex := make(map[int]int)
	// value -> index in original list
	for idx, val := range b {
		reverseIndex[val] = idx
	}
	result := make([]int, len(a))
	for idx, val := range a {
		result[idx] = reverseIndex[val]
	}
	return result
}

func main() {
	a := []int{12, 28, 46, 32, 50}
	b := []int{50, 12, 32, 46, 28}
	fmt.Printf("%+v\n", anagramMappings(a, b))
}
