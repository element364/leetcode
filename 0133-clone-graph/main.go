package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func (this *Node) Print() {
	adj := map[string]bool{}

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, n := range node.Neighbors {
			edgeD := fmt.Sprintf("[%d,%d]", node.Val, n.Val)
			edgeR := fmt.Sprintf("[%d,%d]", n.Val, node.Val)
			_, eD := adj[edgeD]
			_, eR := adj[edgeR]
			if !eD && !eR {
				fmt.Println(edgeD)
				adj[edgeD] = true
				adj[edgeR] = true
				dfs(n)
			}
		}
	}

	dfs(this)
}

func cloneGraph(node *Node) *Node {
	m := map[*Node]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		clone := &Node{Val: node.Val}
		m[node] = clone

		for _, n := range node.Neighbors {
			if _, exists := m[n]; !exists {
				dfs(n)
			}
			clone.Neighbors = append(clone.Neighbors, m[n])
		}

		return clone
	}
	return dfs(node)
}

func main() {
	fmt.Println("Example 1")
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}
	n1.Print()
	c := cloneGraph(n1)
	c.Print()

	fmt.Println("Example 2")
	n1 = &Node{Val: 1}
	n1.Print()
	c = cloneGraph(n1)
	c.Print()

	fmt.Println("Example 3")
	n1 = nil
	c = cloneGraph(n1)
	c.Print()
}
