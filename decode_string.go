package main

import "fmt"
import "strings"
import "strconv"

type istack []int

func (s istack) empty() bool { return len(s) == 0 }
func (s *istack) push(input int) {
	*s = append(*s, input)
}
func (s *istack) pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// returns the index of the last seen key after
func emptyStackScan(input, key string) int {
	s := make(istack, 0)
	ptr := 1
	s.push(0)
	for !s.empty() {
		ch := input[ptr]
		if ch == '[' {
			s.push(ptr)
		} else if ch == ']' {
			s.pop()
		}
		ptr += 1
	}
	return ptr
}

func isNumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func consumeNumber(input string, offset int) (int, int) {
	ptr := offset + 1
	for isNumber(string(input[ptr])) {
		ptr += 1
	}
	window := input[offset:ptr]
	if num, err := strconv.Atoi(input[offset:ptr]); err == nil {
		return num, len(window)
	}
	return -1, -1
}

func decodeString(input string) string {
	var buf strings.Builder
	for ptr := 0; ptr < len(input); ptr += 1 {
		ch := input[ptr]
		// fmt.Printf("--- ptr iter %d (%c)---\n", ptr, ch)
		if count, err := strconv.Atoi(string(ch)); err == nil {
			// check if next char is number too
			ncount, length := consumeNumber(input, ptr)
			count = ncount
			offset := ptr + length
			offsetInput := input[offset:]
			rbIdx := offset + emptyStackScan(offsetInput, "]") - 1
			inside := input[offset+1 : rbIdx]
			if strings.Index(inside, "]") != -1 {
				inside = decodeString(inside)
			}
			// fmt.Printf("{%q < %d} (%q) Found count begin: %d\n", input, rbIdx, inside, count)
			for repeat := 0; repeat < count; repeat += 1 {
				buf.WriteString(inside)
			}
			ptr = rbIdx - 1
		} else if ch != '[' && ch != ']' {
			// fmt.Printf("Added %c to %q\n", ch, buf.String())
			buf.WriteByte(ch)
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(decodeString("100[leetcode]"))
}
