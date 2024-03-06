package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	result := 0

	if root != nil {
		result++

		if root.Left != nil {
			result += countNodes(root.Left)
		}

		if root.Right != nil {
			result += countNodes(root.Right)
		}
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(countNodes(root))

	fmt.Println(countNodes(nil))

	fmt.Println(countNodes(&TreeNode{Val: 1}))
}
