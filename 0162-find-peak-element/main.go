package main

import "fmt"

func findPeakElement(nums []int) int {
	for i, num := range nums {
		if (i == 0 || num > nums[i-1]) && (i == len(nums)-1 || num > nums[i+1]) {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Example 1", findPeakElement(nums))

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println("Example 2", findPeakElement(nums))
}
