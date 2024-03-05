package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	zigZag := false
	for len(queue) > 0 {
		currentLen := len(queue)
		currentItems := []int{}
		for currentLen > 0 {
			current := queue[0]
			queue = queue[1:]

			currentLen--
			currentItems = append(currentItems, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		if len(currentItems) > 0 {
			if zigZag {
				slices.Reverse(currentItems)
			}
			result = append(result, currentItems)
			zigZag = !zigZag
		}
	}
	return result
}

func main() {
	// [3,9,20,null,null,15,7]
	root := TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	// [1]
	// root := TreeNode{Val: 1}

	// [1,2,3,4,null,null,5] -> [[1],[3,2],[4,5]]
	// root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}

	result := zigzagLevelOrder(&root)
	fmt.Println("Result", result)
}
