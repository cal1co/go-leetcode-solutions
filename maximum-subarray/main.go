package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6))
	fmt.Println("2", reflect.DeepEqual(maxSubArray([]int{1}), 1))
	fmt.Println("2", reflect.DeepEqual(maxSubArray([]int{5, 4, -1, 7, 8}), 23))
}

func maxSubArray(nums []int) int {
	out, curr := nums[0], 0
	for _, num := range nums {
		curr += num
		out = max(out, curr)
		if curr < 0 {
			curr = 0
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
