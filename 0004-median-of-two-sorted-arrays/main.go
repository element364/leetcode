package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len((nums2))
	m := total / 2
	p1 := 0
	p2 := 0

	s := make([]int, 0, m)

	for len(s) <= m {
		if p1 < len(nums1) && p2 < len(nums2) {
			if nums1[p1] < nums2[p2] {
				s = append(s, nums1[p1])
				p1++
			} else {
				s = append(s, nums2[p2])
				p2++
			}
		} else if p1 < len(nums1) {
			s = append(s, nums1[p1])
			p1++
		} else if p2 < len(nums2) {
			s = append(s, nums2[p2])
			p2++
		}
	}

	if total&1 == 1 {
		return float64(s[m])
	} else {
		return float64(s[m]+s[m-1]) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("Example 1", findMedianSortedArrays(nums1, nums2)) // 2

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println("Esample 2", findMedianSortedArrays(nums1, nums2)) // 2.5
}
