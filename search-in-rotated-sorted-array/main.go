package main

import "fmt"

func main() {
	fmt.Println("1", search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4)
	fmt.Println("2", search([]int{4, 5, 6, 7, 0, 1, 2}, 3) == -1)
	fmt.Println("3", search([]int{1}, 0) == -1)
}

func search(nums []int, target int) int {
	// get mid point
	// if left <= mid, get what target is closer to and adjust l, r
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}
	return -1
}
