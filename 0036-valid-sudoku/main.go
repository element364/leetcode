package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	const n = 9
	const subStep = 3

	// check rows
	for i := 0; i < n; i++ {
		used := map[byte]bool{}
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				if _, exists := used[board[i][j]]; exists {
					// fmt.Println("Fuck row", i, j)
					return false
				}
				used[board[i][j]] = true
			}
		}
	}
	// check columns
	for j := 0; j < n; j++ {
		used := map[byte]bool{}
		for i := 0; i < n; i++ {
			if board[i][j] != '.' {
				if _, exists := used[board[i][j]]; exists {
					// fmt.Println("Fuck column", i, j)
					return false
				}
				used[board[i][j]] = true
			}
		}
	}
	// check 3x3 squares
	for xs := 0; xs < n; xs += subStep {
		for ys := 0; ys < n; ys += subStep {
			used := map[byte]bool{}
			for i := xs; i < xs+subStep; i++ {
				for j := ys; j < ys+subStep; j++ {
					if board[i][j] != '.' {
						if _, exists := used[board[i][j]]; exists {
							// fmt.Println("Fuck square", i, j)
							return false
						}
						used[board[i][j]] = true
					}
				}
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
