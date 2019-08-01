package main

import "fmt"

func sumWindow(source []int, left, right int) int {
	window := source[left:right]
	fmt.Printf("l=%d, r=%d, s=%+v\n", left, right, window)
	sum := 0
	for _, v := range window {
		sum += v
	}
	return sum
}

func maxSubArray(nums []int) int {
	left := 0
	right := 1
	highest := 0
	for right >= left && right < len(nums)-1 {
		s := sumWindow(nums, left, right)
		for s <= 0 {
			// find start
			left += 1
			right += 1
			s = sumWindow(nums, left, right)
		}
		if s > highest {
			highest = s
		}
		right += 1
	}
	return highest
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, -1}))
	fmt.Println(maxSubArray([]int{0}))
	fmt.Println(maxSubArray([]int{5, -10, 5, 5, 5}))
}
