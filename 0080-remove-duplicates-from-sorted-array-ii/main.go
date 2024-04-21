package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 2
	k := len(nums)
	for i < k {
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			// fmt.Println(i, nums[i])
			for j := i; j < k-1; j++ {
				nums[j] = nums[j+1]
			}
			k--
		} else {
			i++
		}
	}
	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println("Example 1", k, nums)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k = removeDuplicates(nums)
	fmt.Println("Example 2", k, nums)
}
