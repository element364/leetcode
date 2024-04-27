package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	memo := map[string]bool{}

	var dfs func(target, s1, s2 string) bool
	dfs = func(target, s1, s2 string) bool {
		if len(target) == 0 {
			if len(s1) == 0 && len(s2) == 0 {
				return true
			}
			return false
		}
		if len(target) < len(s1) {
			return false
		}

		key := fmt.Sprintf("%s-%s-%s", target, s1, s2)
		if _, exists := memo[key]; !exists {
			// fmt.Println("DFS", target, ",", s1, ",", s2)
			for i := 1; i <= len(s1); i++ {
				if target[:i] == s1[:i] {
					// fmt.Println(" Matches", i, target[:i])
					if dfs(target[i:], s2, s1[i:]) {
						memo[key] = true
						return true
					}
				}
			}
			memo[key] = false
		}

		return memo[key]
	}

	return dfs(s3, s1, s2) || dfs(s3, s2, s1)
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println("Example 1", isInterleave(s1, s2, s3)) // true

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	fmt.Println("Example 2", isInterleave(s1, s2, s3)) // false

	s1 = ""
	s2 = ""
	s3 = ""
	fmt.Println("Example 3", isInterleave(s1, s2, s3)) // true

	s1 = ""
	s2 = "b"
	s3 = "b"
	fmt.Println("Example 4", isInterleave(s1, s2, s3)) // true

	s1 = "a"
	s2 = "b"
	s3 = "a"
	fmt.Println("Example 5", isInterleave(s1, s2, s3)) // false

	s1 = ""
	s2 = "abc"
	s3 = "ab"
	fmt.Println("Example 6", isInterleave(s1, s2, s3)) // false
}
