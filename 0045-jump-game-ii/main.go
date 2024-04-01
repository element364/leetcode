package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	mins := make([]int, len(nums))
	for i := 0; i < len(mins); i++ {
		mins[i] = math.MaxInt
	}
	mins[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := i; j <= i+nums[i] && j < len(nums); j++ {
			mins[j] = min(mins[j], mins[i]+1)
		}
	}
	// fmt.Println("Mins", mins)
	return mins[len(nums)-1]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println("Example 1", jump(nums)) // 2

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println("Example 2", jump(nums)) // 2
}
