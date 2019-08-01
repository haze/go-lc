package main

import "strings"
import "unicode"

func toLowerCase(input string) string {
	var buf strings.Builder
	for _, char := range input {
		buf.WriteRune(unicode.ToLower(char))
	}
	return buf.String()
}
