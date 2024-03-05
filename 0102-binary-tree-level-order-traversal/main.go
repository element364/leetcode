package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for currentLen := len(queue); currentLen > 0; currentLen = len(queue) {
		currentLevel := []int{}
		for i := 0; i < currentLen; i++ {
			current := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}

func main() {
	root := TreeNode{
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
	fmt.Println(levelOrder(&root))

	root = TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrder(&root))

	fmt.Println(levelOrder(nil))
}
