package main

import (
	"fmt"
)

func howManyWords(words string) (ret int) {
	ret = 1
	for _, l := range words {
		if 'A' <= l && l <= 'Z' {
			ret++
		}
	}
	return ret
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Printf("Input is: %s\n", input)
	fmt.Printf("%d", howManyWords(input))
}
