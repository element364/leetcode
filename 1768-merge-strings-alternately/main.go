package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	l1 := len(word1)
	l2 := len(word2)
	maxl := max(l1, l2)
	for i := 0; i < maxl; i++ {
		if i < l1 {
			result += string(word1[i])
		}
		if i < l2 {
			result += string(word2[i])
		}
	}
	return result
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println("Example 1", mergeAlternately(word1, word2)) // apbqcr

	word1 = "ab"
	word2 = "pqrs"
	fmt.Println("Example 2", mergeAlternately(word1, word2)) // apbqrs

	word1 = "abcd"
	word2 = "pq"
	fmt.Println("Example 3", mergeAlternately(word1, word2)) // apbqcd
}
