package main

import "fmt"
import "strings"

func contains(jewels []string, stone string) bool {
	for _, j := range jewels {
		if j == stone {
			return true
		}
	}
	return false
}

func numJewelsInStones(jewelsStr, stonesStr string) int {
	jewels := strings.Split(jewelsStr, "")
	stones := strings.Split(stonesStr, "")
	totalJewels := 0
	for _, stone := range stones {
		if contains(jewels, stone) {
			totalJewels += 1
		}
	}
	return totalJewels
}

func main() {
	jewels := "z"
	stones := "ZZ"
	fmt.Printf("%d", numJewelsInStones(jewels, stones))
}
