package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}}))
	fmt.Println("2", reflect.DeepEqual(combinationSum([]int{2, 3, 5}, 8), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}))
	fmt.Println("2", reflect.DeepEqual(combinationSum([]int{2}, 1), [][]int{}))
}

func combinationSum(candidates []int, target int) [][]int {
	// backtrack O(2^n)
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var backtrack func(idx, currSum int, curr []int)
	backtrack = func(idx, currSum int, curr []int) {
		if currSum == target {
			ans = append(ans, append([]int{}, curr...))
			return
		}
		if currSum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(i, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}

	}
	backtrack(0, 0, curr)
	fmt.Println(ans)
	return ans
}
