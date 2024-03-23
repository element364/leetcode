package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right-left > 1 {
		m := (left + right) / 2
		if nums[m] >= target {
			right = m
		} else {
			left = m
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if target < nums[left] {
		return left
	}
	if target > nums[right] {
		return right + 1
	}
	return right
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 0
	fmt.Println(searchInsert(nums, target))
}
