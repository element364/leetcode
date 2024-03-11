package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func (tree *TreeNode) Print() {
	if tree == nil {
		return
	}
	queue := []TreeNode{*tree}

	for len(queue) > 0 {
		for _, node := range queue {
			fmt.Print(node.Val, " ")
		}
		fmt.Println()
		count := len(queue)
		for i := 0; i < count; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, *current.Left)
			}
			if current.Right != nil {
				queue = append(queue, *current.Right)
			}
		}
	}
}

func main() {
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
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	invertTree(tree1).Print()

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	invertTree(tree2).Print()

	invertTree(nil).Print()
}
