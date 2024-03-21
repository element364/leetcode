package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2

	var head *ListNode = nil
	prev := head

	for p1 != nil && p2 != nil {
		current := &ListNode{}
		if p1.Val < p2.Val {
			current.Val = p1.Val
			p1 = p1.Next
		} else {
			current.Val = p2.Val
			p2 = p2.Next
		}
		if head == nil {
			head = current
			prev = current
		} else {
			prev.Next = current
			prev = prev.Next
		}
	}

	for p1 != nil {
		current := &ListNode{Val: p1.Val}
		if head == nil {
			head = current
			prev = current
		} else {
			prev.Next = current
			prev = prev.Next
		}
		p1 = p1.Next
	}

	for p2 != nil {
		current := &ListNode{Val: p2.Val}
		if head == nil {
			head = current
			prev = current
		} else {
			prev.Next = current
			prev = prev.Next
		}
		p2 = p2.Next
	}

	return head
}

func (this *ListNode) Print() {
	current := this
	fmt.Print("[")
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" ")
		}
		current = current.Next
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Example 1")
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list1.Print()
	list2.Print()
	mergeTwoLists(list1, list2).Print()

	fmt.Println("Example 2")
	list1 = nil
	list2 = nil
	list1.Print()
	list2.Print()
	mergeTwoLists(list1, list2).Print()

	fmt.Println("Example 3")
	list1 = nil
	list2 = &ListNode{Val: 0}
	list1.Print()
	list2.Print()
	mergeTwoLists(list1, list2).Print()
	// list1 = [1,2,4], list2 = [1,3,4]
	// mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode
}
