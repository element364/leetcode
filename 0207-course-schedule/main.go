package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	recStack := make([]bool, numCourses)

	var cycleDfs func(from int) bool
	cycleDfs = func(current int) bool {
		// fmt.Println("DFS", current)
		if !visited[current] {
			visited[current] = true
			recStack[current] = true

			for _, prerequisite := range prerequisites {
				from := prerequisite[0]
				to := prerequisite[1]
				// fmt.Println(" From", from, "To", to)
				if from == current {
					if !visited[to] && cycleDfs(to) {
						return true
					} else if recStack[to] {
						return true
					}
				}
			}
		}

		recStack[current] = false
		return false
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && cycleDfs(i) {
			// fmt.Println("Fuck", i)
			return false
		}
	}

	return true
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))

	numCourses = 2
	prerequisites = [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))

	numCourses = 5
	prerequisites = [][]int{}
	fmt.Println(canFinish(numCourses, prerequisites))
}
