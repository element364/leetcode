package main

import (
	"fmt"
)

type QItem struct {
	Side byte
	Next *QItem
}

func predictPartyVictory(senate string) string {
	var tail *QItem = nil
	var head *QItem = nil
	countsR := 0
	countsD := 0

	for i := 0; i < len(senate); i++ {
		item := &QItem{
			Side: senate[i],
		}
		if head == nil {
			head = item
			tail = item
		} else {
			tail.Next = item
			tail = item
		}

		if senate[i] == 'R' {
			countsR++
		} else {
			countsD++
		}
	}
	tail.Next = head

	curr := head
	for countsR > 0 && countsD > 0 {
		// fmt.Println("Counts radiant", countsR)
		// fmt.Println("Counts dire", countsD)

		// Searching for next enemy
		prevEnemy := curr
		for prevEnemy.Next.Side == curr.Side {
			prevEnemy = prevEnemy.Next
		}

		// Deleting next enemy
		if prevEnemy.Next.Side == 'D' {
			countsD--
		} else {
			countsR--
		}
		prevEnemy.Next = prevEnemy.Next.Next

		curr = curr.Next
	}

	if countsD == 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	fmt.Println("Go")
	fmt.Println("Example 1", predictPartyVictory("RD"))  // Radiant
	fmt.Println("Example 2", predictPartyVictory("RDD")) // Dire
}
