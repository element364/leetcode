package main

import (
	"fmt"
)

func pointToBoustrophedon(x, y, n int) int {
	rtl := x%2 == 1
	if rtl {
		return n*(n-x-1) + y + 1
	}
	return n*(n-x-1) + n - y
}

type Position struct {
	label int
	steps int
}

func boustrophedonToPoint(counter, n int) (int, int) {
	x := n - (counter-1)/n - 1
	rtl := x%2 == n%2
	y := (counter - 1) % n
	if rtl {
		y = n - y - 1
	}
	return x, y
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	// rtl := false
	// board2 := [][]int{}
	// counter := 1
	// for i := n - 1; i >= 0; i-- {
	// 	row := make([]int, n)
	// 	for j := 0; j < n; j++ {
	// 		x, y := boustrophedonToPoint(counter, n)
	// 		if rtl {
	// 			row[n-j-1] = counter
	// 			fmt.Println("Board", i, x, n-j-1, y, counter, pointToBoustrophedon(i, n-j-1, n))
	// 		} else {
	// 			row[j] = counter
	// 			fmt.Println("Board", i, x, j, y, counter, pointToBoustrophedon(i, j, n))
	// 		}
	// 		counter++
	// 	}
	// 	board2 = append([][]int{row}, board2...)
	// 	rtl = !rtl
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Print(board2[i][j], ",", board[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }
	used := map[int]bool{1: true}
	queue := []Position{{
		label: 1,
	}}
	for len(queue) > 0 {
		current := queue[0]
		// fmt.Println("Current", current.label)
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			next := current.label + i
			if next > n*n {
				continue
			}
			x, y := boustrophedonToPoint(next, n)
			if board[x][y] != -1 {
				// fmt.Println("Next", next, "x=", x, "y=", y, "jump", board[x][y])
				next = board[x][y]
			} else {
				// fmt.Println("Next", next, "x=", x, "y=", y)
			}
			if next == n*n {
				return current.steps + 1
			}
			if !used[next] {
				used[next] = true
				queue = append(queue, Position{
					label: next,
					steps: current.steps + 1,
				})
			}
		}
	}
	return -1
}

func main() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	fmt.Println(snakesAndLadders(board))
	board = [][]int{{-1, -1}, {-1, 3}}
	fmt.Println(snakesAndLadders(board))
	board = [][]int{{1, 1, -1}, {1, 1, 1}, {-1, 1, 1}}
	fmt.Println(snakesAndLadders(board))
	board = [][]int{{-1, -1, 19, 10, -1}, {2, -1, -1, 6, -1}, {-1, 17, -1, 19, -1}, {25, -1, 20, -1, -1}, {-1, -1, -1, -1, 15}}
	fmt.Println(snakesAndLadders(board))
}
