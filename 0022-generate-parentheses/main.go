package main

import "fmt"

func generateParenthesis(n int) []string {
	n2 := 2 * n
	result := []string{}

	var dfs func(string, int, int)
	dfs = func(s string, open, close int) {
		if len(s) == n2 {
			result = append(result, s)
			return
		}

		if open < n {
			dfs(s+"(", open+1, close)
		}
		if close < open {
			dfs(s+")", open, close+1)
		}
	}

	dfs("", 0, 0)

	return result
}

func main() {
	fmt.Println("N = 1")
	fmt.Println(generateParenthesis(1))
	fmt.Println("N = 2")
	fmt.Println(generateParenthesis(2))
	fmt.Println("N = 3")
	fmt.Println(generateParenthesis(3))
}
