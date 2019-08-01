package main

import "strings"
import "unicode"
import "fmt"

func isVowel(ichar rune) bool {
	char := unicode.ToLower(ichar)
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func removeVowels(input string) string {
	var buf strings.Builder
	for _, char := range []rune(input) {
		if !isVowel(char) {
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

func main() {
	fmt.Printf("%+v\n", removeVowels("hello,aeiou world!"))
}
