package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x == len(triangle)-1 {
			return triangle[x][y]
		}

		key := fmt.Sprintf("%d-%d", x, y)
		if _, exists := memo[key]; !exists {
			nx := x + 1
			result := triangle[x][y] + dfs(nx, y)

			ny := y + 1
			if ny < len(triangle[nx]) {
				result = min(result, triangle[x][y]+dfs(nx, ny))
			}
			memo[key] = result
		}

		return memo[key]
	}
	return dfs(0, 0)
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println("Example 1", minimumTotal(triangle))

	triangle = [][]int{{-10}}
	fmt.Println("Example 2", minimumTotal(triangle))
}
