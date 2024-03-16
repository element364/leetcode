package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// fmt.Println("  sortedArrayToBST", nums)
	if len(nums) == 0 {
		return nil
	}

	rootIndex := len(nums) / 2
	// fmt.Println("  rootIndex", rootIndex, "Val", nums[rootIndex], "Left", nums[:rootIndex], "Right", nums[rootIndex+1:])
	node := &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
	return node
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		currentLen := len(queue)
		for i := 0; i < currentLen; i++ {
			current := queue[0]
			fmt.Print(current.Val, " ")
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		fmt.Println()
	}
	fmt.Println("=======")
}

func main() {
	fmt.Println("Example 1")
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	printTree(tree)

	fmt.Println("Example 2")
	nums = []int{1, 3}
	tree = sortedArrayToBST(nums)
	printTree(tree)

	fmt.Println("One element array")
	nums = []int{5}
	tree = sortedArrayToBST(nums)
	printTree(tree)

	fmt.Println("Empty array")
	nums = []int{}
	tree = sortedArrayToBST(nums)
	printTree(tree)
}
