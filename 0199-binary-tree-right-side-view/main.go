package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}

	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		currentLen := len(queue)
		lastVal := 0
		for i := 0; i < currentLen; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			lastVal = current.Val
		}

		result = append(result, lastVal)
	}

	return result
}

func main() {
	fmt.Println(rightSideView(nil))

	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(rightSideView(&root))

	root = TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(rightSideView(&root))

	root = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(rightSideView(&root))
}
