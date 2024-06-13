package main

import (
	"fmt"
	"strings"
)

func div(s, d string) bool {
	n := len(s) / len(d)
	repeated := strings.Repeat(d, n)
	return repeated == s
}

func gcdOfStrings(str1 string, str2 string) string {
	result := ""
	for l := 1; l <= len(str1); l++ {
		sub := str1[:l]
		if div(str1, sub) && div(str2, sub) {
			result = sub
		}
	}
	return result
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println("Example 1", gcdOfStrings(str1, str2)) // "ABC"

	str1 = "ABABAB"
	str2 = "ABAB"
	fmt.Println("Example 2", gcdOfStrings(str1, str2)) // "AB"

	str1 = "LEET"
	str2 = "CODE"
	fmt.Println("Example 3", gcdOfStrings(str1, str2)) // ""

	str1 = "AAA"
	str2 = "AAA"
	fmt.Println("Example 4", gcdOfStrings(str1, str2)) // ""
}
