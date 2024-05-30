package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sliceToTree(arr []int, idx int) *TreeNode {
	if idx >= len(arr) {
		return nil
	}
	if arr[idx] == -1 {
		return nil
	}
	return &TreeNode{
		Val:   arr[idx],
		Left:  sliceToTree(arr, 2*idx+1),
		Right: sliceToTree(arr, 2*idx+2),
	}
}

type BSTIterator struct {
	Index   int
	Ordered []int
}

func Constructor(root *TreeNode) BSTIterator {
	result := BSTIterator{
		Index:   -1,
		Ordered: []int{},
	}

	var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}
		rec(node.Left)
		result.Ordered = append(result.Ordered, node.Val)
		rec(node.Right)
	}
	rec(root)

	return result
}

func (this *BSTIterator) Next() int {
	this.Index++
	return this.Ordered[this.Index]
}

func (this *BSTIterator) HasNext() bool {
	return this.Index < len(this.Ordered)-1
}

func main() {
	bSTIterator := Constructor(sliceToTree([]int{7, 3, 15, -1, -1, 9, 20}, 0))
	fmt.Println(bSTIterator.Next())    // return 3
	fmt.Println(bSTIterator.Next())    // return 7
	fmt.Println(bSTIterator.HasNext()) // return True
	fmt.Println(bSTIterator.Next())    // return 9
	fmt.Println(bSTIterator.HasNext()) // return True
	fmt.Println(bSTIterator.Next())    // return 15
	fmt.Println(bSTIterator.HasNext()) // return True
	fmt.Println(bSTIterator.Next())    // return 20
	fmt.Println(bSTIterator.HasNext()) // return False
}
