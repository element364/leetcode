package main

import "fmt"

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	nums = []int{1}

	nums = []int{5, 4, -1, 7, 8}

	fmt.Println(maxSubArray(nums))
}
