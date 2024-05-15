package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0

	var dfs func(node *TreeNode, curr int)
	dfs = func(node *TreeNode, curr int) {
		// fmt.Println("DFS at", node.Val, curr)
		// we are in leaf
		if node.Left == nil && node.Right == nil {
			sum += curr
			return
		}

		if node.Left != nil {
			dfs(node.Left, curr*10+node.Left.Val)
		}
		if node.Right != nil {
			dfs(node.Right, curr*10+node.Right.Val)
		}
	}

	dfs(root, root.Val)

	return sum
}

func sliceToTree(arr []int, i int) *TreeNode {
	// fmt.Println("sliceToTree", arr, i)
	if i >= len(arr) {
		return nil
	}
	return &TreeNode{
		Val:   arr[i],
		Left:  sliceToTree(arr, i*2+1),
		Right: sliceToTree(arr, i*2+2),
	}
}

func main() {
	root := sliceToTree([]int{1, 2, 3}, 0)
	fmt.Println("Example 1", sumNumbers(root))

	root = sliceToTree([]int{4, 9, 0, 5, 1}, 0)
	fmt.Println("Example 2", sumNumbers(root))

	root = sliceToTree([]int{0, 1}, 0)
	fmt.Println("Example 3", sumNumbers(root))
}
