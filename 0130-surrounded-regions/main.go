package main

import "fmt"

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		board[x][y] = 'Y'
		for i := 0; i < len(dx); i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && board[nx][ny] == 'O' {
				dfs(nx, ny)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || i == n-1 || j == 0 || j == m-1) && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

func print(board [][]byte) {
	for _, i := range board {
		for _, j := range i {
			fmt.Print(string(j))
		}
		fmt.Println()
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	fmt.Println("Before")
	print(board)
	solve(board)
	fmt.Println("After")
	print(board)
}
