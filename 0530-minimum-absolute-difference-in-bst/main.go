package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	vals := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		vals = append(vals, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	result := math.MaxInt32
	for i := 1; i < len(vals); i++ {
		result = min(result, vals[i]-vals[i-1])
	}

	return result
}

func main() {
	fmt.Println(getMinimumDifference(nil))

	tree1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(getMinimumDifference(tree1))

	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}
	fmt.Println(getMinimumDifference(tree2))
}
