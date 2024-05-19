package main

import (
	"fmt"
	"math"
)

func lineCoef(a, b []int) float64 {
	coef := float64(b[1]-a[1]) / float64(b[0]-a[0])
	if math.IsInf(coef, 0) {
		coef = math.Inf(1)
	}
	return coef
}

func maxPoints(points [][]int) int {
	result := 0

	for i := 0; i < len(points); i++ {
		coefs := map[float64]int{}

		for j := i + 1; j < len(points); j++ {
			coef := lineCoef(points[i], points[j])
			coefs[coef]++
			result = max(result, coefs[coef])
		}
	}

	return result + 1
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println("Example 1", maxPoints(points)) // 3

	points = [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Println("Example 2", maxPoints(points)) // 4
}
