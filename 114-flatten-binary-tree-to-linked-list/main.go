package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func print(node *TreeNode, level int) {
	if node == nil {
		return
	}
	fmt.Println(node.Val, level)
	print(node.Left, level+1)
	print(node.Right, level+1)
}

func flatten(root *TreeNode) {
	var prev *TreeNode = nil
	// prev := &TreeNode{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		fmt.Println("DFS", node.Val)
		dfs(node.Right)
		dfs(node.Left)
		node.Left = nil
		node.Right = prev
		prev = node
	}

	dfs(root)
}

func main() {
	root := TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}

	fmt.Println("Before")
	print(&root, 0)
	flatten(&root)
	fmt.Println("After")
	print(&root, 0)
}
