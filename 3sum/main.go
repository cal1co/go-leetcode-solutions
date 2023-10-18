package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// i != j, i != k, j != k && nums[i] + nums[j] + nums[j] == 0
	fmt.Println("1", reflect.DeepEqual(threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}}))
	fmt.Println("2", reflect.DeepEqual(threeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}}))
	fmt.Println("3", reflect.DeepEqual(threeSum([]int{0, 1, 1}), []int{}))
}

func threeSum(nums []int) [][]int {
	// sort nums
	// loop through -ve values
	// skip value is duplicate of last seen
	// initiate l and r pointers of next value and last
	// move inwards to get sum of 0

	sort.Ints(nums)
	var out [][]int
	for i, v := range nums {
		if v > 0 {
			break
		}

		if i > 0 && nums[i-1] == v {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum > 0 {
				r--
				continue
			}
			if sum < 0 {
				l++
				continue
			}
			out = append(out, []int{v, nums[l], nums[r]})
			l++
			r--
			for nums[l] == nums[l-1] && l < r {
				l++
			}
		}

	}
	return out
}
