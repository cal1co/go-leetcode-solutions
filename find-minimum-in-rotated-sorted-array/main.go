package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("1", findMin([]int{3, 4, 5, 1, 2}) == 1)
	fmt.Println("2", findMin([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	fmt.Println("3", findMin([]int{11, 13, 15, 17}) == 11)
}

func findMin(nums []int) int {
	// binary search
	// l, r pointers at ends of nums
	// find the midpoint
	// evaluate if nums[l] or nums[r] is smaller and move midpoint to whichever segment is closest
	l, r := 0, len(nums)-1
	minSeen := math.Inf(1)
	for l <= r {
		mid := (l + r) / 2
		minSeen = min(minSeen, float64(nums[mid]))
		if nums[l] > nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return int(minSeen)
}
