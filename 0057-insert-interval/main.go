package main

import (
	"fmt"
	"slices"
)

type Coord struct {
	IsStart bool
	x       int
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// Find insert position
	// insertPosition := 0
	// for ; insertPosition < len(intervals) && intervals[insertPosition][0] < newInterval[0]; insertPosition++ {
	// 	fmt.Println("Interval", intervals[insertPosition], "before", newInterval)
	// }
	// fmt.Println("insert Position", insertPosition)
	// intervals = slices.Insert(intervals, insertPosition, newInterval)
	// fmt.Println("Intervals", intervals)

	// Remove overlaps
	overlaps := []Coord{}
	for _, interval := range intervals {
		overlaps = append(overlaps, Coord{
			x:       interval[0],
			IsStart: true,
		})
		overlaps = append(overlaps, Coord{
			x:       interval[1],
			IsStart: false,
		})
	}
	overlaps = append(overlaps, Coord{
		x:       newInterval[0],
		IsStart: true,
	})
	overlaps = append(overlaps, Coord{
		x:       newInterval[1],
		IsStart: false,
	})
	slices.SortFunc(overlaps, func(a, b Coord) int {
		if a.x == b.x {
			if a.IsStart {
				return -1
			}
			return 1
		}
		return a.x - b.x
	})

	result := [][]int{}
	start := 0
	counter := 0
	for _, coord := range overlaps {
		// fmt.Println("Coord", coord, "count", counter)
		if coord.IsStart {
			if counter == 0 {
				start = coord.x
			}
			counter++
		} else {
			counter--
			if counter == 0 {
				result = append(result, []int{start, coord.x})
			}
		}
	}
	// fmt.Println("Overlaps", overlaps)

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println("Example 1", insert(intervals, newInterval)) // [[1,5],[6,9]]

	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	fmt.Println("Example 2", insert(intervals, newInterval)) // [[1,2],[3,10],[12,16]]

	intervals = [][]int{{1, 2}, {3, 4}, {5, 6}}
	newInterval = []int{7, 8}
	fmt.Println("Example 3", insert(intervals, newInterval)) // [[1 2] [3 4] [5 6] [7 8]]
}
