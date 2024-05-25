package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
	_ "github.com/fatih/color"
)

func trap(height []int) int {
	m := math.MinInt
	for _, n := range height {
		m = max(m, n)
	}

	prevMax := make([]int, len(height))
	prevMax[0] = height[0]
	for x := 1; x < len(height); x++ {
		prevMax[x] = max(prevMax[x-1], height[x])
	}
	// fmt.Println("prevMax", prevMax)

	nextMax := make([]int, len(height))
	nextMax[len(height)-1] = height[len(height)-1]
	for x := len(height) - 2; x >= 0; x-- {
		nextMax[x] = max(nextMax[x+1], height[x])
	}
	// fmt.Println("nextMax", nextMax)

	var check func(x, y int) bool
	check = func(x, y int) bool {
		return x > 0 && prevMax[x-1] >= y && x < len(height)-1 && nextMax[x+1] >= y
		// left := false
		// for x1 := x - 1; x1 >= 0; x1-- {
		// 	if height[x1] >= y {
		// 		left = true
		// 		break
		// 	}
		// }

		// right := false
		// for x2 := x + 1; x2 < len(height); x2++ {
		// 	if height[x2] >= y {
		// 		right = true
		// 		break
		// 	}
		// }
		// return left && right
	}

	result := 0
	for x := 1; x < len(height)-1; x++ {
		pn := prevMax[x-1]
		nn := nextMax[x+1]
		mm := min(pn, nn)
		if mm > height[x] {
			result += mm - height[x]
		}
	}
	return result

	for y := m; y > 0; y-- {
		for x := 0; x < len(height); x++ {
			if y > height[x] {
				if check(x, y) {
					color.Set(color.BgBlue)
					fmt.Print(" ")
					result++
				} else {
					color.Set(color.BgWhite)
					fmt.Print(" ")
				}
			} else {
				color.Set(color.BgBlack)
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return result
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Example 1", trap(height)) // 6

	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println("Example 2", trap(height)) // 9

	height = []int{2, 0, 2}
	fmt.Println("Example 3", trap(height)) // 2
}
