package main

import "fmt"

func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if j == 0 || nums[j-1] != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	fmt.Println("Example 1")
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	fmt.Println("Example 2")
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	fmt.Println("Example 3")
	nums = []int{}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	fmt.Println("Example 4")
	nums = []int{1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
