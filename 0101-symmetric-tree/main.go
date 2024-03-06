package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var dfs func(p, q *TreeNode, level int) bool
	dfs = func(p, q *TreeNode, level int) bool {
		// fmt.Println("Level", level, "Comparing", p.Val, "and", q.Val)
		if p.Val != q.Val {
			return false
		}

		if p.Left != nil && q.Right != nil {
			if !dfs(p.Left, q.Right, level+1) {
				return false
			}
		} else if (p.Left == nil && q.Right != nil) || (p.Left != nil && q.Right == nil) {
			return false
		}

		if p.Right != nil && q.Left != nil {
			if !dfs(p.Right, q.Left, level+1) {
				return false
			}
		} else if (p.Right == nil && q.Left != nil) || (p.Right != nil && q.Left == nil) {
			return false
		}

		return true
	}

	if root != nil {
		return dfs(root, root, 0)
	}

	return true
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(tree))

	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(tree))
}
