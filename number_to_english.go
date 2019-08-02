package main

/*

  1 = One
  7 = Seven
  10 = Ten
  16 = Sixteen
  17 = Seventeen
  117 = One Hundred Seventeen

*/

import (
	"fmt"
	"math"
	"strconv"
)

type ConstructedString struct {
	Segment string
	Next    *ConstructedString
}

func numberOfDigits(num int) int {
	if num < 10 {
		return 1
	}
	return int(math.Ceil(math.Log10(float64(num))))
}

var (
	RootMap map[int]string
)

func init() {
	RootMap = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Four",
		8: "Five",
		9: "Six",
	}
}

func chunkToEnglish(chunk, rank int) string {
	// just for safety
	fmt.Printf("c=%d, r=%d", chunk, rank)
	numDigits := numberOfDigits(chunk)
	if numDigits > 3 {
		panic("chunk may not have over 3 digits")
	} else if numDigits == 1 {
		return RootMap[chunk]
	} else if numDigits == 2 {
		// 10 ..  99

	}
	return ""
}

func rankToEnglish(rank int) (string, bool) {
	switch rank {
	case 1:
		return "Thousand", true
	case 2:
		return "Million", true
	case 3:
		return "Billion", true
	}
	return "", false
}

func chunkifyNumber(num int) []int {
	totalDigits := float64(numberOfDigits(num))
	// left to right, index 0 = lowest part of string
	numStr := strconv.Itoa(num)
	result := make([]int, int(math.Ceil(totalDigits/3)))
	counter := 0
	for numStr != "" {
		size := len(numStr)
		var bite int
		if size >= 3 {
			bite = 3
		} else {
			bite = size
		}
		chunk := numStr[size-bite:]
		numStr = numStr[:size-bite]
		chunkInt, err := strconv.Atoi(chunk)
		if err != nil {
			continue
		}
		result[counter] = chunkInt
		counter += 1
	}
	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func numberToWords(num int) string {
	chunks := chunkifyNumber(num)
	dummy := ConstructedString{
		Segment: "",
		Next:    nil,
	}
	// rank
	// value
	fmt.Printf("%+v\n", chunks)
	for index, chunk := range chunks {
		fmt.Println(chunkToEnglish(chunk, index))
	}
	return dummy.Segment
}

func main() {
	fmt.Println(numberToWords(1))
	fmt.Println(numberToWords(15))
	fmt.Println(numberToWords(153))
	fmt.Println(numberToWords(25333))
	fmt.Println(numberToWords(12345678))
}
