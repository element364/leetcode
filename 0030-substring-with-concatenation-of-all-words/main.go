package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	n := len(words)
	k := len(words[0])
	l := k * n

	wordCounts := map[string]int{}
	for _, word := range words {
		wordCounts[word]++
	}

	var check func(start int) bool
	check = func(start int) bool {
		used := map[string]int{}
		found := 0
		for i := start; i < start+l; i += k {
			sub := s[i : i+k]
			if wordCounts[sub]-used[sub] > 0 {
				used[sub]++
				found++
			} else {
				break
			}
		}
		return found == n
	}

	result := []int{}
	for i := 0; i < len(s)-l+1; i++ {
		if check(i) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println("Example 1", findSubstring(s, words)) // [0, 9]

	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "word"}
	fmt.Println("Example 2", findSubstring(s, words)) // []

	s = "barfoofoobarthefoobarman"
	words = []string{"bar", "foo", "the"}
	fmt.Println("Example 3", findSubstring(s, words)) // [6, 9, 12]

	s = "qwerty"
	words = []string{"y"}
	fmt.Println("Example 4", findSubstring(s, words))
}
