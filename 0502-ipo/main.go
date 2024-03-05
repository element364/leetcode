package main

import (
	"fmt"
	"slices"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type Project struct {
	capital int
	profit  int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// greedy approach
	n := len(profits)
	projects := make([]Project, n)
	for i := 0; i < n; i++ {
		projects[i] = Project{
			capital: capital[i],
			profit:  profits[i],
		}
	}

	slices.SortFunc(projects, func(a, b Project) int {
		return a.capital - b.capital
	})

	heap := binaryheap.NewWith(func(a, b interface{}) int {
		return b.(Project).profit - a.(Project).profit
	})

	i := 0
	for k > 0 {
		// Adding all realizable projects to heap
		for i < n && projects[i].capital <= w {
			heap.Push(projects[i])
			i++
		}

		if heap.Empty() {
			break
		}

		// Realizing most profitable project
		top, _ := heap.Pop()
		w += top.(Project).profit
		k--
	}

	// fmt.Println("Projects", projects)

	return w
}

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))

	k = 3
	w = 0
	profits = []int{1, 2, 3}
	capital = []int{0, 1, 2}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))

	k = 1
	w = 0
	profits = []int{1, 2, 3}
	capital = []int{1, 1, 2}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))
}
