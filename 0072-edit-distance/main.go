package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	memo := map[string]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j < 0 {
			return i + 1
		}
		if i < 0 {
			return j + 1
		}

		key := fmt.Sprintf("%d-%d", i, j)
		if _, exists := memo[key]; !exists {
			if word1[i] == word2[j] {
				memo[key] = dfs(i-1, j-1)
			} else {
				del := 1 + dfs(i-1, j)
				insert := 1 + dfs(i, j-1)
				replace := 1 + dfs(i-1, j-1)

				memo[key] = min(del, insert, replace)
			}
		}

		return memo[key]
	}
	return dfs(len(word1)-1, len(word2)-1)
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println("Example 1", minDistance(word1, word2)) // 3

	word1 = "intention"
	word2 = "execution"
	fmt.Println("Example 2", minDistance(word1, word2)) // 5
}
