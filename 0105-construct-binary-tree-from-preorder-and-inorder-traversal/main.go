package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) Print() {
	fmt.Println("Print")
	queque := []*TreeNode{}
	queque = append(queque, this)
	for len(queque) > 0 {
		level := []int{}
		for i := len(queque); i > 0; i-- {
			current := queque[0]
			level = append(level, current.Val)
			queque = queque[1:]
			if current.Left != nil {
				queque = append(queque, current.Left)
			}
			if current.Right != nil {
				queque = append(queque, current.Right)
			}
		}
		fmt.Println(level)
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// fmt.Println("buildTree")
	// fmt.Println("preorder", preorder)
	// fmt.Println("inorder", inorder)

	if len(preorder) == 0 {
		return nil
	}

	idx := slices.Index(inorder, preorder[0])

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func main() {
	fmt.Println("Example 1")
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	tree.Print()

	fmt.Println("Example 2")
	preorder = []int{-1}
	inorder = []int{-1}
	tree = buildTree(preorder, inorder)
	tree.Print()
}
