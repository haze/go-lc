package main

import "fmt"

func fixedPoint(a []int) int {
	for idx, val := range a {
		if val < 0 {
			continue
		}
		if idx == val {
			return idx
		}
	}
	return -1
}

func main() {
	a := []int{-10, -5, 0, 3, 7}
	fmt.Printf("%d\n", fixedPoint(a))
}
