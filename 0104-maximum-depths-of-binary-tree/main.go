package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	result := 0

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		result = max(result, depth)

		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}

	if root != nil {
		dfs(root, 1)
	}

	return result
}

func main() {
	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(maxDepth(tree1))

	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(maxDepth(tree2))

	fmt.Println(maxDepth(nil))
}
