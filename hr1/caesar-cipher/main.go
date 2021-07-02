package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var n, k int
	var input string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &k)

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	ret := ""
	for _, r := range input {
		isUpper := false
		if 'A' <= r && r <= 'Z' {
			isUpper = true
		}
		rotated := string(rotate(r, k, alphabet))
		if isUpper {
			ret += strings.ToUpper(rotated)
		} else {
			ret += rotated
		}
	}
	fmt.Printf("%s\n", ret)
}

func rotate(r rune, k int, alphabet []rune) rune {
	idx := strings.IndexRune(string(alphabet), unicode.ToLower(r))
	if idx < 0 {
		return r
	}
	idx = (idx + k) % len(alphabet)
	return alphabet[idx]
}
