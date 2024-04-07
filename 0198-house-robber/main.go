package main

import "fmt"

func rob(nums []int) int {
	memo := map[string]int{}

	var dfs func(p int, sum int) int
	dfs = func(p int, sum int) int {
		key := fmt.Sprintf("%d-%d", p, sum)
		value, exists := memo[key]
		if exists {
			// fmt.Println("For", key, "=", value)
			return value
		}

		if p >= len(nums) {
			memo[key] = sum
			return sum
		}

		best := sum
		best = max(best, dfs(p+2, sum+nums[p]))
		best = max(best, dfs(p+1, sum))
		memo[key] = best
		return best
	}

	return dfs(0, 0)
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Example 1", rob(nums)) // 4

	nums = []int{2, 7, 9, 3, 1}
	fmt.Println("Example 2", rob(nums)) // 12

	nums = []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	fmt.Println("Example 3", rob(nums)) // 4173
}
