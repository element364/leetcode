package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	// fmt.Println("minPathSum", grid)
	memo := map[string]int{}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		// fmt.Println("We are in", x, y)

		if x == len(grid)-1 && y == len(grid[x])-1 {
			// fmt.Println("[+] We finished, returning", grid[x][y])
			return grid[x][y]
		}

		key := fmt.Sprintf("%d-%d", x, y)
		if _, exists := memo[key]; !exists {
			result := math.MaxInt
			if x < len(grid)-1 {
				// fmt.Println(" We can go bot", x+1, y)
				result = min(result, grid[x][y]+dfs(x+1, y))
			}
			if y < len(grid[x])-1 {
				// fmt.Println(" We can go right", x, y+1)
				result = min(result, grid[x][y]+dfs(x, y+1))
			}
			memo[key] = result
		}
		return memo[key]
	}
	result := dfs(0, 0)
	return result
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println("Example 1", minPathSum(grid))

	grid = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Example 2", minPathSum(grid))
}
