package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	hash1 := map[int]bool{}
	for _, num1 := range nums1 {
		hash1[num1] = true
	}

	hash2 := map[int]bool{}
	for _, num2 := range nums2 {
		hash2[num2] = true
	}

	notPresentInNums2 := []int{}
	for k := range hash1 {
		if _, exists := hash2[k]; !exists {
			notPresentInNums2 = append(notPresentInNums2, k)
		}
	}

	notPresentInNums1 := []int{}
	for k := range hash2 {
		if _, exists := hash1[k]; !exists {
			notPresentInNums1 = append(notPresentInNums1, k)
		}
	}

	return [][]int{
		notPresentInNums2,
		notPresentInNums1,
	}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println("Example 1", findDifference(nums1, nums2)) // [[1,3],[4,6]]

	nums1 = []int{1, 2, 3, 3}
	nums2 = []int{1, 1, 2, 2}
	fmt.Println("Example 2", findDifference(nums1, nums2)) // [[3],[]]
}
