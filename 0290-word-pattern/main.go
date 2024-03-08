package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	charToWord := make(map[rune]string)
	usedWords := make(map[string]bool)
	for i, c := range pattern {
		if word, exists := charToWord[c]; exists {
			if word != words[i] {
				return false
			}
		} else {
			if usedWords[words[i]] {
				return false
			}
			usedWords[words[i]] = true
			charToWord[c] = words[i]
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s)) // true

	pattern = "abba"
	s = "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s)) // false

	pattern = "aaaa"
	s = "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s)) // false

	pattern = "abba"
	s = "dog dog dog dog"
	fmt.Println(wordPattern(pattern, s)) // false

	pattern = "aaa"
	s = "aa aa aa aa"
	fmt.Println(wordPattern(pattern, s)) // false
}
