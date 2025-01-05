package main

import "fmt"

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	// fmt.Println(" Ping", t)
	// fmt.Println(this.queue)
	this.queue = append(this.queue, t)

	result := 0
	idx := len(this.queue) - 1
	limit := t - 3000
	for idx >= 0 && this.queue[idx] >= limit {
		result++
		idx--
	}
	return result
}

func main() {
	fmt.Println("Go")
	rc := Constructor()
	fmt.Println(rc.Ping(1))    // 1
	fmt.Println(rc.Ping(100))  // 2
	fmt.Println(rc.Ping(3001)) // 3
	fmt.Println(rc.Ping(3002)) // 3
}
