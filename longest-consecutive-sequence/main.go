package main

import "fmt"

func main() {
	fmt.Println("1 should be 4:", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println("2 should be 9:", longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println("3 should be 1:", longestConsecutive([]int{0}))
}

func longestConsecutive(nums []int) int {

	uniqueNums := map[int]bool{}

	for _, num := range nums {
		if _, ok := uniqueNums[num]; !ok {
			uniqueNums[num] = true
		}
	}

	out := 0

	for _, num := range nums {
		if _, ok := uniqueNums[num-1]; !ok {
			tmp := 0
			for {
				if _, ok := uniqueNums[num+tmp]; ok {
					tmp++
					out = max(out, tmp)
				} else {
					break
				}
			}
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
