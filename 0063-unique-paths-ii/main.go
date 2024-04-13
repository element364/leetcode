package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := map[string]int{}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if obstacleGrid[x][y] == 1 {
			return 0
		}

		if x == len(obstacleGrid)-1 && y == len(obstacleGrid[x])-1 {
			return 1
		}

		key := fmt.Sprintf("%d-%d", x, y)
		if _, exists := memo[key]; !exists {

			result := 0
			// fmt.Println("(", x, ",", y, ")")
			nx := x + 1
			if nx < len(obstacleGrid) {
				// fmt.Println(" can go bot (", nx, ",", y, ")")
				result += dfs(nx, y)
			}
			ny := y + 1
			if ny < len(obstacleGrid[x]) {
				// fmt.Println(" can go left (", x, ",", ny, ")")
				result += dfs(x, ny)
			}

			memo[key] = result
		}

		return memo[key]
	}

	return dfs(0, 0)
}

func main() {
	obstacleGrid := [][]int{{0}}
	fmt.Println("Example 00", uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{1}}
	fmt.Println("Example 01", uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{1, 0}}
	fmt.Println("Example 03", uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println("Example 1", uniquePathsWithObstacles(obstacleGrid))

	obstacleGrid = [][]int{{0, 1}, {0, 0}}
	fmt.Println("Example 2", uniquePathsWithObstacles(obstacleGrid))
}
