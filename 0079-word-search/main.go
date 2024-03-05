package main

import "fmt"

func exist(board [][]byte, word string) bool {
	// fmt.Println(word)

	// var print = func() {
	// 	for i := 0; i < len(board); i++ {
	// 		for j := 0; j < len(board[i]); j++ {
	// 			fmt.Printf(string(board[i][j]))
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, p int) bool {
		if p == len(word) {
			return true
		}
		for k := 0; k < 4; k++ {
			x := i + dx[k]
			y := j + dy[k]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[i]) && board[x][y] == word[p] {
				board[x][y] = ' '
				if dfs(x, y, p+1) {
					return true
				}
				board[x][y] = word[p]
			}
		}
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				board[i][j] = ' '
				if dfs(i, j, 1) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	// fmt.Println(exist(board, word))

	word = "SEE"
	fmt.Println(exist(board, word))

	// word = "ABCB"
	// fmt.Println(exist(board, word))
}
