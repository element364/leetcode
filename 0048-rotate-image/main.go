package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	for k := 0; k < n/2; k++ {
		sx := k
		sy := k
		fx := n - k - 1
		fy := n - k - 1

		// snake rotate

		for turns := 0; turns < n-2*k-1; turns++ {
			x := sx
			y := sy
			direction := 0
			// snake := []int{}
			prev := matrix[x][y]
			for direction < 4 {
				nx := x + dx[direction]
				ny := y + dy[direction]
				if nx >= sx && nx <= fx && ny >= sy && ny <= fy {
					// snake = append(snake, matrix[x][y])
					tmp := matrix[x][y]
					matrix[x][y] = prev
					prev = tmp
					x = nx
					y = ny
				} else {
					direction++
				}
			}
			matrix[x][y] = prev
			// fmt.Println("Snake", snake)
			// for i := 0; i < n; i++ {
			// 	for j := 0; j < n; j++ {
			// 		fmt.Print(matrix[i][j], " ")
			// 	}
			// 	fmt.Println()
			// }
		}
	}
}

func main() {
	fmt.Println("Minus 1")
	matrix := [][]int{}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	fmt.Println("Zero")
	matrix = [][]int{
		{1},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	fmt.Println("One")
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix = [][]int{
		{4, 1, 2},
		{7, 5, 3},
		{8, 9, 6},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	fmt.Println("Two")
	matrix = [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)

	fmt.Println("Three")
	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}
