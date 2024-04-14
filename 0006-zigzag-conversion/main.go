package main

import "fmt"

func convert(s string, numRows int) string {
	rows := make([][]string, numRows)
	for i := range rows {
		rows[i] = make([]string, len(s))
	}

	x := 0
	y := 0

	dx := []int{1, -1}
	dy := []int{0, 1}
	d := 0

	for _, c := range s {
		rows[x][y] = string(c)
		if x == numRows-1 {
			d = 1
		} else if x == 0 {
			d = 0
		}
		x += dx[d]
		x = max(x, 0)
		y += dy[d]
	}

	result := ""
	for _, row := range rows {
		for _, c := range row {
			if c != "" {
				result += c
			}
		}
	}

	return result
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println("Example 1", convert(s, numRows)) // PAHNAPLSIIGYIR

	s = "PAYPALISHIRING"
	numRows = 4
	fmt.Println("Example 2", convert(s, numRows)) // PINALSIGYAHRPI

	s = "A"
	numRows = 1
	fmt.Println("Example 3", convert(s, numRows)) // A

	s = "AB"
	numRows = 1
	fmt.Println("Example 4", convert(s, numRows)) // A
}
