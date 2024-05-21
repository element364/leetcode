package main

import "fmt"

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

func (head *ListNode) Print() {
	fmt.Print("[")
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, ", ")
	}
	fmt.Println("]")
}

func sortList(head *ListNode) *ListNode {
	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next {
		for ptr := curr.Next; ptr != nil; ptr = ptr.Next {
			if ptr.Val < curr.Val {
				t := curr.Val
				curr.Val = ptr.Val
				ptr.Val = t
			}
		}
	}
	return head
}

func main() {
	head := sliceToList([]int{4, 2, 1, 3})
	fmt.Println("Example 1")
	head.Print()
	head = sortList(head)
	head.Print()

	head = sliceToList([]int{-1, 5, 3, 4, 0})
	fmt.Println("Example 2")
	head.Print()
	head = sortList(head)
	head.Print()

	head = sliceToList([]int{})
	fmt.Println("Example 3")
	head.Print()
	head = sortList(head)
	head.Print()
}
