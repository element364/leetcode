package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) Print() {
	fmt.Print("[")
	for curr := this; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, ", ")
	}
	fmt.Println("]")
}

func sliceToList(s []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	for _, item := range s {
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

func rotateRight(head *ListNode, k int) *ListNode {
	tail := head
	count := 1
	for ; tail != nil && tail.Next != nil; tail = tail.Next {
		count++
	}
	// fmt.Println("Count", count)

	k = k % count
	for i := 0; i < k; i++ {
		prev := head
		for prev != nil && prev.Next != nil && prev.Next.Next != nil {
			prev = prev.Next
		}

		if tail != nil && prev != nil {
			tail.Next = head
			prev.Next = nil
			head = tail
			tail = prev
		}
	}
	return head
}

func main() {
	fmt.Println("Example 1")
	head := sliceToList([]int{1, 2, 3, 4, 5})
	k := 2
	head.Print()
	head = rotateRight(head, k)
	head.Print()

	fmt.Println("Example 2")
	head = sliceToList([]int{0, 1, 2})
	k = 4
	head.Print()
	head = rotateRight(head, k)
	head.Print()
}
