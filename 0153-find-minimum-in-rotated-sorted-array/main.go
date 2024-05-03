package main

import "fmt"

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m
		} else {
			r = m
		}
	}
	return min(nums[l], nums[r])
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println("Example 1", findMin(nums))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Example 2", findMin(nums))

	nums = []int{11, 13, 15, 17}
	fmt.Println("Example 3", findMin(nums))
}
