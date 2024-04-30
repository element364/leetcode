package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	result := len(nums) + 1
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			result = min(result, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if result == len(nums)+1 {
		return 0
	}

	return result
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println("Example 1", minSubArrayLen(target, nums)) // 2

	target = 4
	nums = []int{1, 4, 4}
	fmt.Println("Example 2", minSubArrayLen(target, nums)) // 1

	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println("Example 3", minSubArrayLen(target, nums)) // 0
}
