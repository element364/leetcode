package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	bank := 0
	total := 0
	start := 0

	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		bank += diff
		total += diff
		if bank < 0 {
			start = i + 1
			bank = 0
		}
	}

	if total >= 0 {
		return start
	}

	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println("Example 1", canCompleteCircuit(gas, cost)) // 3

	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println("Example 2", canCompleteCircuit(gas, cost)) // -1

	gas = []int{5, 1, 2, 3, 4}
	cost = []int{4, 4, 1, 5, 1}
	fmt.Println("Example 3", canCompleteCircuit(gas, cost)) // 4
}
