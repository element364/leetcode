package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n < 1 {
		return 0
	}
	m := len(matrix[0])

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				memo[i][j] = 1
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				memo[i][j] = min(memo[i-1][j], memo[i][j-1], memo[i-1][j-1]) + 1
			}
		}
	}

	// fmt.Println("After")
	// for i := 0; i < n; i++ {
	// 	fmt.Println(memo[i])
	// }

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result = max(result, memo[i][j])
		}
	}

	return result * result
}

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println("Example 1", maximalSquare(matrix)) // 4

	matrix = [][]byte{{'0', '1'}, {'1', '0'}}
	fmt.Println("Example 2", maximalSquare(matrix)) // 1

	matrix = [][]byte{{'0'}}
	fmt.Println("Example 3", maximalSquare(matrix)) // 0

	matrix = [][]byte{
		{'1', '0', '1', '0', '0', '1', '1', '1', '0'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'0', '1', '1', '0', '0', '1', '0', '0', '1'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '1', '1', '0', '1', '0', '0', '0', '1'},
		{'0', '1', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '0', '0', '1', '1', '1', '0', '0', '0'},
	}
	fmt.Println("Example 4", maximalSquare(matrix)) // 0
}
