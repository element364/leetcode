package main

import "fmt"

func numIslands(grid [][]byte) int {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = 'x'
		for k := range dx {
			nx := x + dx[k]
			ny := y + dy[k]
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == '1' {
				dfs(nx, ny)
			}
		}
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Example 1", numIslands(grid)) // 1

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Example 2", numIslands(grid)) // 3
}
