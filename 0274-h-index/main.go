package main

import (
	"fmt"
	"sort"
)

// The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
func hIndex(citations []int) int {
	sort.Ints(citations)
	// fmt.Println(citations)
	for i, count := range citations {
		left := len(citations) - i
		// fmt.Println("article", i, "has", count, "citations")
		// fmt.Println(" ", i, "articles with <", count, "cit")
		// fmt.Println(" ", left, "articles with at least", count, "cit")
		if left <= count {
			return left
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println("Example 1", hIndex(citations)) // 3

	citations = []int{1, 3, 1}
	fmt.Println("Example 2", hIndex(citations)) // 1

	citations = []int{100}
	fmt.Println("Example 3", hIndex(citations)) // 1
}
