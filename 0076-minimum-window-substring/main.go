package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	tCounts := map[byte]int{}
	for i := range t {
		tCounts[t[i]]++
	}
	tTotal := len(tCounts)

	result := ""

	l := 0
	minL := math.MaxInt
	subCounts := map[byte]int{}
	subTotal := 0

	for r := 0; r < len(s); r++ {
		c := s[r]

		subCounts[c]++
		if subCounts[c] == tCounts[c] {
			subTotal++
		}

		for subTotal == tTotal {
			if r-l < minL {
				minL = r - l
				result = s[l : r+1]
			}

			subCounts[s[l]]--
			if subCounts[s[l]] < tCounts[s[l]] {
				subTotal--
			}

			l++
		}
	}

	return result
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("Example 1", minWindow(s, t)) // "BANC"

	s = "a"
	t = "a"
	fmt.Println("Example 2", minWindow(s, t)) // "a"

	s = "a"
	t = "aa"
	fmt.Println("Example 3", minWindow(s, t)) // ""

	s = "aa"
	t = "aa"
	fmt.Println("Example 4", minWindow(s, t)) // "aa"
}
