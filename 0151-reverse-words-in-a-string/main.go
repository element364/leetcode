package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := regexp.MustCompile(`\s+`).Split(strings.Trim(s, " "), -1)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

func main() {
	s := "the sky is blue"
	fmt.Println("Example 1", reverseWords(s))

	s = "  hello world  "
	fmt.Println("Example 2", reverseWords(s))

	s = "a good   example"
	fmt.Println("Example 3", reverseWords(s))
}
