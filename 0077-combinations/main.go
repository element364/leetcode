package main

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	current := []int{}

	var rec func(s, k int)
	rec = func(s, k int) {
		if k == 0 {
			copied := make([]int, len(current))
			copy(copied, current)
			result = append(result, copied)
			return
		}
		for i := s; i <= n; i++ {
			current = append(current, i)
			rec(i+1, k-1)
			current = current[:len(current)-1]
		}
	}
	rec(1, k)

	return result
}

func main() {
	n := 1
	k := 1
	result := combine(n, k)
	fmt.Println("Result", result)
}
