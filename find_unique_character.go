package main

import "fmt"
import "sort"

func findUniqChar(s string) int {
	charmap := make(map[rune][]int, 0)
	for idx, ch := range s {
		indicies, exists := charmap[ch]
		if exists {
			charmap[ch] = append(indicies, idx)
		} else {
			charmap[ch] = []int{idx}
		}
	}
	unique := make([]struct {
		char     rune
		indicies []int
	}, 0)
	for char, indicies := range charmap {
		if len(indicies) == 1 {
			unique = append(unique, struct {
				char     rune
				indicies []int
			}{char: char, indicies: indicies})
		}
	}
	if len(unique) == 0 {
		return -1
	}
	sort.SliceStable(unique, func(i, j int) bool {
		return unique[i].indicies[0] < unique[j].indicies[0]
	})
	return unique[0].indicies[0]
}

func main() {
	fmt.Println(findUniqChar("leetcode"))
	fmt.Println(findUniqChar("loveleetcode"))
}
