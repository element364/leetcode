package main

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}

	start := nums[0]
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i-1]+1 != nums[i] {
			if start == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}

	return result
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("Example 1", summaryRanges(nums))

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println("Example 2", summaryRanges(nums))
}
