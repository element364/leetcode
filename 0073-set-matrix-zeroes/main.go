package main

import "fmt"

type Point struct {
	x int
	y int
}

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n < 1 {
		return
	}
	m := len(matrix[0])

	zerosPositions := []Point{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zerosPositions = append(zerosPositions, Point{i, j})
			}
		}
	}

	for _, zeroPosition := range zerosPositions {
		for i := 0; i < n; i++ {
			matrix[i][zeroPosition.y] = 0
		}
		for j := 0; j < m; j++ {
			matrix[zeroPosition.x][j] = 0
		}
	}
}

func main() {
	fmt.Println("One")
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)

	fmt.Println("Two")
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)

	fmt.Println("Three")
	matrix = [][]int{{1, 2, 3}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)

	fmt.Println("Four")
	matrix = [][]int{}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
