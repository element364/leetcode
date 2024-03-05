package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)

	result := []int{}
	if n == 0 {
		return result
	}

	m := len(matrix[0])
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	x := 0
	y := 0
	d := 0

	for len(result) < n*m {
		result = append(result, matrix[x][y])
		matrix[x][y] = -48463

		for turns := 0; turns < 4; turns++ {
			nx := x + dx[d]
			ny := y + dy[d]
			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[nx]) && matrix[nx][ny] != -48463 {
				// fmt.Println(x, y, "->", nx, ny)
				x = nx
				y = ny
				break
			} else {
				d = (d + 1) % 4
			}
		}
	}
	return result
}

func main() {
	matrix := [][]int{}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{{1, 2, 3, 4, 5}}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
