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

func Count(this *ListNode) int {
	count := 0
	for curr := this; curr != nil; curr = curr.Next {
		count++
	}
	return count
}

func Split(this *ListNode, m int) (*ListNode, *ListNode) {
	curr := this
	for i := 1; i < m; i++ {
		curr = curr.Next
	}

	r := curr.Next
	curr.Next = nil
	return this, r
}

func sortList(head *ListNode) *ListNode {
	count := Count(head)
	if count < 2 {
		return head
	}

	mid := count / 2
	l, r := Split(head, mid)
	l = sortList(l)
	r = sortList(r)

	// Merge
	var res *ListNode
	var cur *ListNode
	for lp, rp := l, r; lp != nil || rp != nil; {
		n := &ListNode{}
		if lp != nil && rp != nil {
			if lp.Val < rp.Val {
				n.Val = lp.Val
				lp = lp.Next
			} else {
				n.Val = rp.Val
				rp = rp.Next
			}
		} else if lp != nil {
			n.Val = lp.Val
			lp = lp.Next
		} else if rp != nil {
			n.Val = rp.Val
			rp = rp.Next
		}

		if res == nil {
			res = n
			cur = res
		} else {
			cur.Next = n
			cur = n
		}
	}
	return res
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
