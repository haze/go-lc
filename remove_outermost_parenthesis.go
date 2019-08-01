package main

import "fmt"
import "strings"

type sstack []string

func (s *sstack) empty() bool { return len(*s) == 0 }
func (s *sstack) push(item string) {
	*s = append(*s, item)
}
func (s *sstack) pop() string {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func removeOutermostParenthesis(input string) string {
	var buf strings.Builder
	parens := make(sstack, 0)
	for _, ch := range input {
		if ch == '(' {
			if !parens.empty() {
				// not first (
				buf.WriteRune(ch)
			}
			parens.push(string(ch))
		} else if ch == ')' {
			// we encountered an end one, pop
			// we could also write a check here but
			// the problem says input is valid
			parens.pop()
			if !parens.empty() {
				buf.WriteRune(ch)
			}
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(removeOutermostParenthesis("(()())(())"))
}
