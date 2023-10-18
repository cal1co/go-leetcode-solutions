package main

import "fmt"

func main() {
	fmt.Println("1", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println("2", maxArea([]int{1, 1}) == 1)
}

func maxArea(height []int) int {
	// move inwards from sides depending on what's smaller l or r.
	// keep out var and calculate area while l < r

	l, r := 0, len(height)-1
	var out int
	for l < r {
		// calculate container area
		out = max(out, (r-l)*min(height[l], height[r]))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
