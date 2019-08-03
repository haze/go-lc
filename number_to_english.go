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
	"strings"
	"unicode"
)

type ConstructedString struct {
	Segment string
	Next    *ConstructedString
}

func numberOfDigits(inum int) int {
	num := inum
	if num < 2 && num > 0 {
		num = 2
	}
	if num%10 == 0 {
		num += 1
	}
	return int(math.Ceil(math.Log10(float64(num))))
}

var (
	RootMap map[int]string
	TeenMap map[int]string
	TensMap map[int]string
	RankMap map[int]string
)

func init() {
	RootMap = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	TeenMap = map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	TensMap = map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	RankMap = map[int]string{
		1: "Thousand",
		2: "Million",
		3: "Billion",
	}
}

func splitTensAndOnes(num int) (int, int) {
	// lol
	s := []rune(strconv.Itoa(num))
	// fmt.Printf("s=%q\n", s)
	if len(s) != 2 {
		panic(fmt.Sprintf("should not happen, i have %d", num))
	}
	left := s[0]
	right := s[1]
	ln, err := strconv.Atoi(string(left))
	rn, errr := strconv.Atoi(string(right))
	if err != nil || errr != nil {
		panic("should not happen")
	}
	return ln, rn
}

func splitHundredsAndTens(num int) (int, int) {
	// lol
	s := []rune(strconv.Itoa(num))
	if len(s) != 3 {
		panic("should not happen")
	}
	left := s[0]
	right := s[1:]
	ln, err := strconv.Atoi(string(left))
	rn, errr := strconv.Atoi(string(right))
	if err != nil || errr != nil {
		panic("should not happen")
	}
	return ln, rn
}

func tensFormat(chunk int, builder *strings.Builder) {
	if chunk > 9 && chunk < 20 {
		builder.WriteString(TeenMap[chunk])
	} else {
		tens, ones := splitTensAndOnes(chunk)
		if ones != 0 {
			builder.WriteString(TensMap[tens] + RootMap[ones])
		} else {
			builder.WriteString(TensMap[tens])
		}
	}
}

func chunkToEnglish(chunk, rank int) string {
	var buf strings.Builder
	chunkToEnglishRec(chunk, rank, &buf)
	return buf.String()
}

func chunkToEnglishRec(chunk, rank int, buf *strings.Builder) {
	// fmt.Printf("processing chunk: %d, rank: %d\n", chunk, rank)
	// just for safety
	numDigits := numberOfDigits(chunk)
	// fmt.Printf("c=%d, r=%d, nd=%d\n", chunk, rank, numDigits)
	if numDigits > 3 {
		panic("chunk may not have over 3 digits")
	} else if numDigits == 1 {
		buf.WriteString(RootMap[chunk])
	} else if numDigits == 2 {
		tensFormat(chunk, buf)
	} else if numDigits == 3 {
		root, tens := splitHundredsAndTens(chunk)
		buf.WriteString(RootMap[root] + "Hundred")
		if tens > 9 {
			tensFormat(tens, buf)
		} else {
			chunkToEnglishRec(tens, 0, buf)
		}
	}
	if rankStr, ok := RankMap[rank]; ok && chunk != 0 {
		buf.WriteString(rankStr)
	}
}

func chunkifyNumber(num int) []int {
	ndigits := numberOfDigits(num)
	totalDigits := float64(ndigits)
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
	if num == 0 {
		return "Zero"
	}
	chunks := chunkifyNumber(num)
	fmt.Printf("%+v\n", chunks)
	var buf strings.Builder
	for index, chunk := range chunks {
		buf.WriteString(chunkToEnglish(chunk, len(chunks)-index-1))
	}
	return strings.TrimSpace(addSpaces(buf.String()))
}

func addSpaces(source string) string {
	var buf strings.Builder
	for _, ch := range []rune(source) {
		if unicode.IsUpper(ch) {
			buf.WriteRune(' ')
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

func main() {
	fmt.Println(numberToWords(0))
	fmt.Println(numberToWords(1))
	fmt.Println(numberToWords(1000000))
	fmt.Println(numberToWords(666666))
	// fmt.Printf(numberToWords(666666))
	// for n := 0; n < 1000000; n++ {
	// 	fmt.Printf("%d: %q\n", n, numberToWords(n))
	// }
}
