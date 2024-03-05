package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	for _, preprerequisite := range prerequisites {
		from := preprerequisite[0]
		to := preprerequisite[1]
		edges[from] = append(edges[from], to)
	}

	// fmt.Println("Edges", edges)

	result := []int{}

	visited := make([]bool, numCourses)
	recStack := make([]bool, numCourses)

	var cycleDfs func(current int) bool
	cycleDfs = func(current int) bool {
		visited[current] = true
		recStack[current] = true

		for _, to := range edges[current] {
			if !visited[to] && cycleDfs(to) {
				return true
			} else if recStack[to] {
				return true
			}
		}

		recStack[current] = false
		result = append(result, current)
		return false
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && cycleDfs(i) {
			return []int{}
		}
	}

	return result
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(findOrder(numCourses, prerequisites))

	numCourses = 4
	prerequisites = [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourses, prerequisites))

	numCourses = 1
	prerequisites = [][]int{}
	fmt.Println(findOrder(numCourses, prerequisites))

	numCourses = 2
	prerequisites = [][]int{{0, 1}, {1, 0}}
	fmt.Println(findOrder(numCourses, prerequisites))
}
