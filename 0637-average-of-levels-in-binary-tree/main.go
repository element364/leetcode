package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentLen := len(queue)
		var sum float64 = 0
		for i := 0; i < currentLen; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			sum += float64(current.Val)
		}
		result = append(result, sum/float64(currentLen))
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
	fmt.Println(averageOfLevels(&root))

	root = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 20,
		},
	}
	fmt.Println(averageOfLevels(&root))
}
