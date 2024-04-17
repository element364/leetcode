package main

import "fmt"

func search(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println("Example 1", search(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println("Example 2", search(nums, target))

	nums = []int{1}
	target = 0

	fmt.Println("Example 3", search(nums, target))
}
