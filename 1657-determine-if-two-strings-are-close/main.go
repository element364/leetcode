package main

import (
	"fmt"
	"slices"
)

type Count struct {
	Char  rune
	Count int
}

func getCounts(s string) []Count {
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}
	// fmt.Println("Counts", counts)

	result := []Count{}
	for k, v := range counts {
		result = append(result, Count{
			Char:  k,
			Count: v,
		})
	}
	slices.SortFunc(result, func(a, b Count) int {
		return b.Count - a.Count
	})
	// fmt.Println("Result", result)

	return result
}

func closeStrings(word1 string, word2 string) bool {
	counts1 := getCounts(word1)
	counts2 := getCounts(word2)

	if len(counts1) != len(counts2) {
		return false
	}

	hash := map[rune]bool{}
	for _, count := range counts1 {
		hash[count.Char] = true
	}
	for _, count := range counts2 {
		if _, exists := hash[count.Char]; !exists {
			return false
		}
	}

	for i := range counts1 {
		if counts1[i].Count != counts2[i].Count {
			return false
		}
	}

	return true
}

func main() {
	word1 := "abc"
	word2 := "bca"
	fmt.Println("Example 1", closeStrings(word1, word2)) // true

	word1 = "a"
	word2 = "aa"
	fmt.Println("Example 2", closeStrings(word1, word2)) // false

	word1 = "cabbba"
	word2 = "abbccc"
	fmt.Println("Example 3", closeStrings(word1, word2)) // true

	word1 = "uau"
	word2 = "ssx"
	fmt.Println("Example 4", closeStrings(word1, word2)) // false
}
