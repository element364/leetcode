package main

import "fmt"

func gameOfLife(board [][]int) {
	n := len(board)
	m := len(board[0])

	result := make([][]int, n)
	for x := 0; x < n; x++ {
		result[x] = make([]int, m)
	}

	dx := []int{1, 0, -1, 0, 1, 1, -1, -1}
	dy := []int{0, 1, 0, -1, 1, -1, 1, -1}

	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			liveCells := 0
			for k := 0; k < 8; k++ {
				nx := x + dx[k]
				ny := y + dy[k]
				if nx >= 0 && nx < n && ny >= 0 && ny < m {
					liveCells += board[nx][ny]
				}
			}

			if board[x][y] == 1 && liveCells < 2 {
				result[x][y] = 0
			} else if board[x][y] == 1 && (liveCells == 2 || liveCells == 3) {
				result[x][y] = 1
			} else if board[x][y] == 1 && liveCells > 3 {
				result[x][y] = 0
			} else if board[x][y] == 0 && liveCells == 3 {
				result[x][y] = 1
			}
		}
	}

	copy(board, result)
}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println("Example 1", board)
	gameOfLife(board)
	fmt.Println(board)

	board = [][]int{{1, 1}, {1, 0}}
	fmt.Println("Example 1", board)
	gameOfLife(board)
	fmt.Println(board)
}
