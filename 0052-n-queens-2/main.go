package main

import "fmt"

func totalNQueens(n int) int {
	total := 0

	// Encode board as one dimension arrya
	// [2, 0, 3, 1]
	// [' ', ' ', '♖', ' ']
	// [' ', ' ', ' ', '♖']
	// ['♖', ' ', ' ', ' ']
	// [' ', '♖', ' ', ' ']
	// 1, 3, 0, 2
	// [' ', ' ', '♖', ' ']
	// ['♖', ' ', ' ', ' ']
	// [' ', ' ', ' ', '♖']
	// [' ', '♖', ' ', ' ']
	current := make([]int, n)
	usedV := make([]bool, n)
	usedD1 := make([]bool, 2*n-1)
	usedD2 := make([]bool, 2*n-1)

	var dfs func(position int)
	dfs = func(position int) {
		if position == n {
			fmt.Println("DFS", current)
			total++
			return
		}
		for i := 0; i < n; i++ {
			if !usedV[i] && !usedD1[position+i] && !usedD2[position-i+n-1] {
				usedV[i] = true
				usedD1[position+i] = true
				usedD2[position-i+n-1] = true
				current[position] = i
				dfs(position + 1)
				usedD2[position-i+n-1] = false
				usedD1[position+i] = false
				usedV[i] = false
			}
		}
	}

	dfs(0)

	return total
}

func main() {
	n := 4
	fmt.Println(totalNQueens(n))

	n = 1
	fmt.Println(totalNQueens(n))
}
