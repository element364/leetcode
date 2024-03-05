package main

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}

	N := len(nums)
	current := make([]int, N)
	used := make([]bool, N)
	var dfs func(p int)
	dfs = func(p int) {
		if p == N {
			copied := make([]int, len(current))
			copy(copied, current)
			result = append(result, copied)
		}
		for i := 0; i < N; i++ {
			if !used[i] {
				used[i] = true
				current[p] = nums[i]
				dfs(p + 1)
				used[i] = false
			}
		}
	}
	dfs(0)

	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

	nums = []int{0, 1}
	fmt.Println(permute(nums))

	nums = []int{1}
	fmt.Println(permute(nums))
}
