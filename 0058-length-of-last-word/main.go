package main

import (
	"fmt"
	"regexp"
	"strings"
)

func lengthOfLastWord(s string) int {
	spaceRE := regexp.MustCompile(`\s+`)
	words := spaceRE.Split(strings.Trim(s, " "), -1)
	// fmt.Println("WORDS", strings.Join(words, "|"))
	return len(words[len(words)-1])
}

func main() {
	s := "Hello World"
	fmt.Println("Example 1", lengthOfLastWord(s))

	s = "   fly me   to   the moon  "
	fmt.Println("Example 2", lengthOfLastWord(s))

	s = "luffy is still joyboy"
	fmt.Println("Example 3", lengthOfLastWord(s))
}
