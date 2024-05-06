package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(arr []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	for _, item := range arr {
		node := &ListNode{Val: item}
		if head == nil {
			head = node
			curr = head
		} else {
			curr.Next = node
			curr = curr.Next
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

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	var next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev, head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var leftPtr *ListNode
	var rightPtr *ListNode

	var reversed *ListNode

	i := 1
	var prev *ListNode
	for curr := head; curr != nil; {
		// fmt.Println("Kek", i, curr.Val)
		if i == left {
			leftPtr = prev
			reversed = curr
		}
		i++
		prev = curr
		curr = curr.Next
		if i == right+1 {
			rightPtr = curr
			prev.Next = nil
		}
	}

	if leftPtr != nil {
		leftPtr.Next = nil
	}
	// fmt.Println("LeftPtr")
	// leftPtr.Print()
	// fmt.Println("Reversed")
	// reversed.Print()
	reversedH, reversedT := reverse(reversed)
	// reversed.Print()
	reversedT.Next = rightPtr
	if leftPtr != nil {
		leftPtr.Next = reversedH
	} else {
		return reversedH
	}

	// fmt.Println("RightPtr")
	// rightPtr.Print()

	return head
}

func main() {
	fmt.Println("Example 1")
	head := SliceToList([]int{1, 2, 3, 4, 5})
	head.Print()
	left := 2
	right := 4
	reversed := reverseBetween(head, left, right)
	reversed.Print()

	fmt.Println("Example 2")
	head = SliceToList([]int{5})
	head.Print()
	left = 1
	right = 1
	reversed = reverseBetween(head, left, right)
	reversed.Print()

	fmt.Println("Example 3")
	head = SliceToList([]int{3, 5})
	head.Print()
	left = 1
	right = 2
	reversed = reverseBetween(head, left, right)
	reversed.Print()
}
