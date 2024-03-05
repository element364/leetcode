package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/binaryheap"
)

type Pair struct {
	i   int
	j   int
	sum int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := [][]int{}

	visited := map[string]bool{"0-0": true}
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		return a.(Pair).sum - b.(Pair).sum
	})
	heap.Push(Pair{
		i:   0,
		j:   0,
		sum: nums1[0] + nums2[0],
	})

	for len(result) < k && !heap.Empty() {
		top, _ := heap.Pop()
		current := top.(Pair)
		result = append(result, []int{nums1[current.i], nums2[current.j]})

		nextI := current.i + 1
		idxI := string(nextI) + "." + string(current.j)
		if nextI < len(nums1) && !visited[idxI] {
			heap.Push(Pair{
				i:   nextI,
				j:   current.j,
				sum: nums1[nextI] + nums2[current.j],
			})
			visited[idxI] = true
		}

		nextJ := current.j + 1
		idxJ := string(current.i) + "." + string(nextJ)
		if nextJ < len(nums2) && !visited[idxJ] {
			heap.Push(Pair{
				i:   current.i,
				j:   nextJ,
				sum: nums1[current.i] + nums2[nextJ],
			})
			visited[idxJ] = true
		}
	}

	return result
}

func main() {
	fmt.Println(string(25) + "." + string(123))
	fmt.Println(fmt.Sprintf("%d", 25) + "." + fmt.Sprintf("%d", 123))

	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 2, 4, 5, 6}
	nums2 = []int{3, 5, 7, 9}
	k = 20
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
