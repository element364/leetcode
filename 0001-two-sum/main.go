package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Example 1", twoSum(nums, target)) // [0,1]

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println("Example 2", twoSum(nums, target)) // [1,2]

	nums = []int{3, 3}
	target = 6
	fmt.Println("Example 3", twoSum(nums, target)) // [0,1]
}
