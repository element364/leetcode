package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return reflect.DeepEqual(p, q)
}

func main() {
	fmt.Println(isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
	))

	fmt.Println(isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	))

	fmt.Println(isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
	))
}
