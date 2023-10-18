package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1", maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println("2", maxProfit([]int{7, 6, 4, 3, 1}) == 0)
}

func maxProfit(prices []int) int {
	// track min number seen and calculate profit off of that for each subsequent day
	minSeen := math.Inf(1)
	var out float64
	for _, val := range prices {
		minSeen = min(float64(val), minSeen)
		out = max(out, float64(val)-minSeen)
	}
	return int(out)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
