package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func findKthLargest(nums []int, k int) int {
	heap := binaryheap.NewWith(utils.IntComparator(b, a))

	for _, num := range nums {
		if heap.Size() > k {
			heap.Pop()
			heap.
		}
		heap.Push(num)
	}

	fmt.Println("Values", heap.Values())
	result, _ := heap.Peek()

	return result.(int)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
