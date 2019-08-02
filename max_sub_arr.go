package main

import "fmt"
import "math"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	table := make([]int, len(nums))
	ptr := 0
	for ptr < len(nums) {
		if ptr-1 < 0 {
			table[ptr] = nums[ptr]
		} else {
			n := nums[ptr]
			table[ptr] = max(n+table[ptr-1], n)
		}
		ptr += 1
	}
	highest := math.Inf(-1)
	for _, v := range table {
		vf := float64(v)
		if vf > highest {
			highest = vf
		}
	}
	return int(highest)
}

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}
