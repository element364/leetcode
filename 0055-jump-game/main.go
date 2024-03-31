package main

import "fmt"

func canJump(nums []int) bool {
	limit := 0
	for i := 0; i < len(nums); i++ {
		if i > limit {
			return false
		}
		limit = max(limit, i+nums[i])
	}
	return true
}

func main() {
	fmt.Println("Example 2")
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums)) // true

	fmt.Println("Example 2")
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums)) // false
}
