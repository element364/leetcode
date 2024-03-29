package main

import "fmt"

func threeSum(nums []int) [][]int {
	h := map[int]int{}
	for _, num := range nums {
		h[num]++
	}
	// fmt.Println("H", h)
	uniq := []int{}
	for k := range h {
		uniq = append(uniq, k)
	}
	// fmt.Println("Uniq", uniq)

	result := [][]int{}
	for i := 0; i < len(uniq); i++ {
		a := uniq[i]
		for j := i; j < len(uniq); j++ {
			b := uniq[j]
			c := -a - b
			if c >= a && c >= b {
				counts := map[int]int{}
				counts[a]++
				counts[b]++
				counts[c]++

				canAdd := true
				for k, v := range counts {
					if v > h[k] {
						canAdd = false
						break
					}
				}
				if !canAdd {
					continue
				}

				result = append(result, []int{a, b, c})
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Example 1")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

	fmt.Println("Example 2")
	nums = []int{0, 1, 1}
	fmt.Println(threeSum(nums))

	fmt.Println("Example 3")
	nums = []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
