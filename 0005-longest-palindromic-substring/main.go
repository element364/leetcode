package main

import "fmt"

func longestPalindrome(s string) string {
	result := ""

	var dfs func(b, e int)
	dfs = func(b, e int) {
		// fmt.Println("DFS", b, e, s[b:e])
		if e-b > len(result) {
			result = s[b:e]
		}
		// can exapnd by 2 ?
		if b > 0 && e < len(s) && s[b-1] == s[e] {
			// fmt.Println("Expanding", s[b:e], "by 2")
			dfs(b-1, e+1)
		}

		// can expand by 1 ?
		if e-b == 1 && e < len(s) && s[b] == s[e] {
			// fmt.Println("Expanding", s[b:e], "by 1")
			dfs(b, e+1)
		}
	}

	for s := range s {
		dfs(s, s+1)
	}

	return result
}

func main() {
	s := "babad"
	fmt.Println("Example 1", longestPalindrome(s)) // bab

	s = "cbbd"
	fmt.Println("Example 2", longestPalindrome(s)) // bb

	s = "bb"
	fmt.Println("Example 3", longestPalindrome(s)) // bb

	s = "aba"
	fmt.Println("Example 4", longestPalindrome(s)) // aba

	s = "aacabdkacaa"
	fmt.Println("Example 5", longestPalindrome(s)) // "aca"
}
