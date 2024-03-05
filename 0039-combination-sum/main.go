package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	current := []int{}

	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			copied := make([]int, len(current))
			copy(copied, current)
			result = append(result, copied)
			return
		}

		if target-candidates[start] >= 0 {
			current = append(current, candidates[start])
			dfs(start, target-candidates[start])
			current = current[:len(current)-1]
		}
		if start+1 < len(candidates) {
			dfs(start+1, target)
		}
	}
	dfs(0, target)

	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))

	candidates = []int{2, 3, 5}
	target = 8
	fmt.Println(combinationSum(candidates, target))

	candidates = []int{2}
	target = 1
	fmt.Println(combinationSum(candidates, target))
}
