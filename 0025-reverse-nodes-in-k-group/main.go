package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(s []int) *ListNode {
	var head *ListNode
	var curr *ListNode

	for _, val := range s {
		node := &ListNode{Val: val}
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

func (this *ListNode) Print(end *ListNode) {
	fmt.Print("[")
	for curr := this; curr != end; curr = curr.Next {
		fmt.Print(curr.Val, ", ")
	}
	fmt.Println("]")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	values := []int{}

	i := 0
	for curr := head; curr != nil; curr = curr.Next {
		values = append(values, curr.Val)
		if i == 0 {
			start = curr
			values = []int{curr.Val}
		}
		i++
		if i == k {
			for len(values) > 0 {
				start.Val = values[len(values)-1]
				values = values[:len(values)-1]
				start = start.Next
			}
			i = 0
		}
	}

	return head
}

func main() {
	fmt.Println("Example 1")
	head := sliceToList([]int{1, 2, 3, 4, 5})
	k := 2
	head.Print(nil)
	head = reverseKGroup(head, k)
	head.Print(nil) // [2,1,4,3,5]

	fmt.Println("Example 2")
	head = sliceToList([]int{1, 2, 3, 4, 5})
	k = 3
	head.Print(nil)
	head = reverseKGroup(head, k)
	head.Print(nil) // [2,1,4,3,5]
}
