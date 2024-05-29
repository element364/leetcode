package main

import (
	"fmt"
	"slices"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})

	arrows := 1

	x := points[0][1]
	for _, p := range points {
		x1 := p[0]
		x2 := p[1]
		if !(x1 <= x && x <= x2) {
			// Не попали в шарик
			x = x2
			arrows++
		}
	}

	return arrows
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println("Example 1", findMinArrowShots(points)) // 2

	points = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println("Example 2", findMinArrowShots(points)) // 4

	points = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println("Example 3", findMinArrowShots(points)) // 2
}
