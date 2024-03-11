package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	found := false

	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node.Left == nil && node.Right == nil && sum == targetSum {
			found = true
		}
		if node.Left != nil && !found {
			dfs(node.Left, sum+node.Left.Val)
		}

		if node.Right != nil && !found {
			dfs(node.Right, sum+node.Right.Val)
		}
	}

	if root != nil {
		dfs(root, root.Val)
	}

	return found
}

func main() {
	tree1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(hasPathSum(tree1, 22)) // true

	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(hasPathSum(tree2, 5)) // false

	fmt.Println(hasPathSum(nil, 0)) // false
}
