package main

import (
	"fmt"
	"slices"
	"unicode"
)

func isVowel(c rune) bool {
	c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func reverseVowels(s string) string {
	vowels := []rune{}
	for _, c := range s {
		if isVowel(c) {
			vowels = append(vowels, c)
		}
	}
	slices.Reverse(vowels)

	result := ""
	v := 0
	for _, c := range s {
		if isVowel(c) {
			result += string(vowels[v])
			v++
		} else {
			result += string(c)
		}
	}

	return result
}

func main() {
	fmt.Println("Example 1", reverseVowels("hello"))    // "holle"
	fmt.Println("Example 2", reverseVowels("leetcode")) // "leotcede"
	fmt.Println("Example 3", reverseVowels("aA"))       // "Aa"
}
