package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(canJump([]int{2, 3, 1, 1, 4}), true))
	fmt.Println("1", reflect.DeepEqual(canJump([]int{3, 2, 1, 0, 4}), false))
}

func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
