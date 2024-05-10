package main

import (
	"fmt"
)

func maxSubarraySumCircular(nums []int) int {
	sum := nums[0]
	minSum := nums[0]
	curMin := nums[0]
	maxSum := nums[0]
	curMax := nums[0]

	for i := 1; i < len(nums); i++ {
		curMax = max(curMax+nums[i], nums[i])
		maxSum = max(maxSum, curMax)

		curMin = min(curMin+nums[i], nums[i])
		minSum = min(minSum, curMin)

		sum += nums[i]
	}

	if sum == minSum {
		return maxSum
	}

	return max(maxSum, sum-minSum)
}

func main() {
	nums := []int{1, -2, 3, -2}
	fmt.Println("Example 1", maxSubarraySumCircular(nums)) // 3

	nums = []int{5, -3, 5}
	fmt.Println("Example 2", maxSubarraySumCircular(nums)) // 10

	nums = []int{-3, -2, -3}
	fmt.Println("Example 3", maxSubarraySumCircular(nums)) // -2
}
