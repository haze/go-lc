package main

import "fmt"
import "strings"

func defangIPaddr(addr string) string {
	var buf strings.Builder
	for _, r := range []rune(addr) {
		if r == '.' {
			buf.WriteString("[.]")
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func main() {
	fmt.Printf("%q\n", defangIPaddr("1.1.1.1"))
}
