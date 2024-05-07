package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(arr []int) *ListNode {
	var head *ListNode
	var curr *ListNode

	for _, item := range arr {
		node := &ListNode{Val: item}
		if head == nil {
			head = node
			curr = head
		} else {
			curr.Next = node
			curr = node
		}
	}

	return head
}

func (this *ListNode) Print() {
	fmt.Print("[")
	for curr := this; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, ",")
	}
	fmt.Println("]")
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	removed := false

	var rec func(node *ListNode) int
	rec = func(node *ListNode) int {
		if node == nil {
			return 0
		}
		nFromEnd := rec(node.Next)
		if nFromEnd == n {
			if node.Next != nil {
				node.Next = node.Next.Next
				removed = true
			}
		}
		return nFromEnd + 1
	}
	rec(head)

	if !removed {
		return head.Next
	}

	return head
}

func main() {
	fmt.Println("Example 1")
	head := sliceToList([]int{1, 2, 3, 4, 5})
	n := 2
	head.Print()
	head = removeNthFromEnd(head, n)
	head.Print() // [1,2,3,5]

	fmt.Println("Example 2")
	head = sliceToList([]int{1})
	n = 1
	head.Print()
	head = removeNthFromEnd(head, n)
	head.Print() // []

	fmt.Println("Example 3")
	head = sliceToList([]int{1, 2})
	n = 1
	head.Print()
	head = removeNthFromEnd(head, n)
	head.Print() // [1]

	fmt.Println("Example 4")
	head = sliceToList([]int{1, 2})
	n = 2
	head.Print()
	head = removeNthFromEnd(head, n)
	head.Print() // [2]
}
