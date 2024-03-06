package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max *int) bool
	dfs = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}

		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}

		return dfs(node.Left, min, &node.Val) && dfs(node.Right, &node.Val, max)
	}

	return dfs(root, nil, nil)
}

func main() {
	tree1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isValidBST(tree1))

	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isValidBST(tree2))

	tree3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(isValidBST(tree3))

	fmt.Println(isValidBST(nil))
}
