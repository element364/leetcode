package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	result := make([]int, n)
	for j := 0; j < n; j++ {
		nj := (j + k) % n
		result[nj] = nums[j]
	}
	copy(nums, result)
}

func main() {
	fmt.Println("Example 1")
	nums := []int{1, 2, 3, 4, 5, 6, 7} // [5,6,7,1,2,3,4]
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	fmt.Println("Example 2")
	nums = []int{-1, -100, 3, 99} // [3,99,-1,-100]
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}
