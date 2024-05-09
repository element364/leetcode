package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) Print() {
	fmt.Println("Print")
	queue := []*TreeNode{this}
	for len(queue) > 0 {
		level := []int{}
		for currLen := len(queue); currLen > 0; currLen-- {
			current := queue[0]
			level = append(level, current.Val)
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		fmt.Println(level)
	}
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	pivot := 0
	for pivot < n && inorder[pivot] != postorder[n-1] {
		pivot++
	}

	return &TreeNode{
		Val:   postorder[n-1],
		Left:  buildTree(inorder[:pivot], postorder[:pivot]),
		Right: buildTree(inorder[pivot+1:], postorder[pivot:n-1]),
	}
}

func main() {
	fmt.Println("Example 1")
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := buildTree(inorder, postorder)
	tree.Print()

	fmt.Println("Example 2")
	inorder = []int{-1}
	postorder = []int{-1}
	tree = buildTree(inorder, postorder)
	tree.Print()
}
