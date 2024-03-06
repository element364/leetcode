package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var result int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if k-count > 0 {
			dfs(node.Left)
		}
		if k-count > 0 {
			result = node.Val
			count++
			dfs(node.Right)
		}
	}

	dfs(root)

	return result
}

func main() {
	fmt.Println(kthSmallest(nil, 2))

	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(kthSmallest(tree1, 1))

	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(kthSmallest(tree2, 3))
}
