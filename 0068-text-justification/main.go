package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	curr := []string{}
	lettersCount := 0

	for _, word := range words {
		if len(word)+len(curr)+lettersCount > maxWidth {
			for i := 0; i < maxWidth-lettersCount; i++ {
				curr[i%max(1, len(curr)-1)] += " "
			}
			result = append(result, strings.Join(curr, ""))
			curr = curr[:0]
			lettersCount = 0
		}
		curr = append(curr, word)
		lettersCount += len(word)
	}

	last := strings.Join(curr, " ")
	for len(last) < maxWidth {
		last += " "
	}
	result = append(result, last)

	return result
}

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	fmt.Println("Example 1", fullJustify(words, maxWidth))

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	fmt.Println("Example 2", fullJustify(words, maxWidth))

	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	fmt.Println("Example 3", fullJustify(words, maxWidth))
}
