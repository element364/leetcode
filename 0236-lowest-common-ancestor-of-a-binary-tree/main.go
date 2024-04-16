package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueItem struct {
	node *TreeNode
	prev *QueueItem
}

func bfs(node *TreeNode, target *TreeNode) *QueueItem {
	// fmt.Println("BFS")
	queue := []*QueueItem{{
		node: node,
	}}
	for len(queue) > 0 {
		// toPrint := []int{}
		currentLen := len(queue)
		for i := 0; i < currentLen; i++ {
			current := queue[0]
			if current.node == target {
				// fmt.Println("Found target", target.Val)
				return current
			}
			// toPrint = append(toPrint, current.node.Val)
			queue = queue[1:]
			if current.node.Left != nil {
				queue = append(queue, &QueueItem{
					node: current.node.Left,
					prev: current,
				})
			}
			if current.node.Right != nil {
				queue = append(queue, &QueueItem{
					node: current.node.Right,
					prev: current,
				})
			}
		}
		// fmt.Println(toPrint)
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := bfs(root, p)
	pathQ := bfs(root, q)
	found := map[*TreeNode]bool{}
	// fmt.Println("Path P")
	for pathP != nil {
		// fmt.Println(pathP.node.Val)
		found[pathP.node] = true
		pathP = pathP.prev
	}

	// fmt.Println("Path Q")
	for pathQ != nil {
		if _, exists := found[pathQ.node]; exists {
			return pathQ.node
		}
		// fmt.Println(pathQ.node.Val)
		pathQ = pathQ.prev
	}

	return nil
}

func main() {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	root := &TreeNode{
		Val:   3,
		Left:  p,
		Right: q,
	}
	fmt.Println("Example 1", lowestCommonAncestor(root, p, q).Val)

	q = &TreeNode{
		Val: 4,
	}
	p = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: q,
		},
	}
	root = &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println("Example 2", lowestCommonAncestor(root, p, q).Val)

	q = &TreeNode{
		Val: 2,
	}
	p = &TreeNode{
		Val:  1,
		Left: q,
	}
	root = p
	fmt.Println("Example 3", lowestCommonAncestor(root, p, q).Val)
}
