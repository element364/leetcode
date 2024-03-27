package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var before *ListNode
	var beforeCurr *ListNode
	var after *ListNode
	var afterCurr *ListNode

	curr := head
	for curr != nil {
		node := &ListNode{
			Val: curr.Val,
		}
		if node.Val < x {
			if beforeCurr == nil {
				before = node
				beforeCurr = before
			} else {
				beforeCurr.Next = node
				beforeCurr = node
			}
		} else {
			if afterCurr == nil {
				after = node
				afterCurr = after
			} else {
				afterCurr.Next = node
				afterCurr = node
			}
		}
		curr = curr.Next
	}

	if beforeCurr != nil {
		beforeCurr.Next = after
		return before
	} else {
		return after
	}
}

func sliceToList(a []int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for _, item := range a {
		node := &ListNode{
			Val: item,
		}
		if head == nil {
			head = node
			prev = node
		} else {
			prev.Next = node
			prev = node
		}
	}
	return head
}

func (this *ListNode) print() {
	curr := this
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Example 1")
	arr := []int{1, 4, 3, 2, 5, 2}
	head := sliceToList(arr)
	x := 3
	head.print()
	partition(head, x).print()

	fmt.Println("Example 2")
	arr = []int{2, 1}
	head = sliceToList(arr)
	x = 2
	head.print()
	partition(head, x).print()

	fmt.Println("Example 3")
	arr = []int{}
	head = sliceToList(arr)
	x = 0
	head.print()
	partition(head, x).print()

	fmt.Println("Example 4")
	arr = []int{1}
	head = sliceToList(arr)
	x = 0
	head.print()
	partition(head, x).print()
}
