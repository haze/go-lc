package main

import "fmt"
import "strconv"
import "math"

func isArmstrong(number int) bool {
	numStr := strconv.Itoa(number)
	digits := len(numStr)
	cubesum := 0.0
	for _, char := range numStr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		cubesum += math.Pow(float64(digit), float64(digits))
	}
	return float64(number) == cubesum
}

func main() {
	fmt.Printf("%t\n", isArmstrong(153))
}
