package main

import "fmt"

func binSearch(nums *[]int, target int) (int, int) {
	l := 0
	r := len(*nums) - 1
	for r-l > 1 {
		m := (l + r) / 2
		if (*nums)[m] > target {
			r = m
		} else {
			l = m
		}
	}
	return l, r
}

func searchMatrix(matrix [][]int, target int) bool {
	M := len(matrix[0]) - 1
	rowLast := make([]int, len(matrix))
	for i, row := range matrix {
		rowLast[i] = row[M]
	}

	r1, r2 := binSearch(&rowLast, target)

	var row *[]int
	if target <= rowLast[r1] {
		row = &matrix[r1]
	} else {
		row = &matrix[r2]
	}

	l, r := binSearch(row, target)
	return (*row)[l] == target || (*row)[r] == target
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 3))

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 13))
}
