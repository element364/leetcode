package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 1

	var canAdd = func(l, r int) bool {
		// fmt.Println("[ ] canAdd", l, r)
		sub := s[l:r]

		counts := map[rune]int{}
		for _, c := range sub {
			if counts[c] > 0 {
				// fmt.Println("[-] canAdd", sub)
				return false
			}
			counts[c]++
		}

		// fmt.Println("[+] canAdd", sub)
		return true
	}
	result := 0

	for l < len(s) {
		for r <= len(s) && canAdd(l, r) {
			result = max(result, r-l)
			r++
		}
		l++
	}

	return result
}

func main() {
	s := "abc"
	fmt.Println("0:0", s[0:0])
	fmt.Println("0:1", s[0:1])
	fmt.Println("0:2", s[0:2])
	fmt.Println("0:3", s[0:3])

	s = "abcabcbb"
	fmt.Println("Example 1", lengthOfLongestSubstring(s)) // 3

	s = "bbbbb"
	fmt.Println("Example 2", lengthOfLongestSubstring(s)) // 1

	s = "pwwkew"
	fmt.Println("Example 3", lengthOfLongestSubstring(s)) // 3

	s = " "
	fmt.Println("Example 4", lengthOfLongestSubstring(s)) // 1
}
