package main

import (
	"fmt"
	"slices"
)

type Point struct {
	x int
	t int
}

func merge(intervals [][]int) [][]int {
	points := []Point{}
	for _, interval := range intervals {
		start := interval[0]
		finish := interval[1]
		points = append(points, Point{
			x: start,
			t: 1,
		})
		points = append(points, Point{
			x: finish,
			t: -1,
		})
	}
	slices.SortFunc(points, func(a, b Point) int {
		if a.x == b.x {
			return b.t - a.t
		}
		return a.x - b.x
	})

	result := [][]int{}
	count := 0
	start := -1
	for _, point := range points {
		count += point.t
		if start == -1 {
			start = point.x
		}
		if count == 0 {
			result = append(result, []int{start, point.x})
			start = -1
		}
	}
	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Example 1", merge(intervals))

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println("Example 2", merge(intervals))
}
