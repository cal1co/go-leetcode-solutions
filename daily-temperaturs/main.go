package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0}))
	fmt.Println("2", reflect.DeepEqual(dailyTemperatures([]int{30, 40, 50, 60}), []int{1, 1, 1, 0}))
	fmt.Println("3", reflect.DeepEqual(dailyTemperatures([]int{30, 60, 90}), []int{1, 1, 0}))
}

func dailyTemperatures(temperatures []int) []int {
	// init slice with length of temperatures
	// loop backwards
	// keep track of another variable which is i + 1
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		j := i + 1
		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			if res[j] <= 0 {
				break
			}
			j += res[j]
		}
		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			res[i] = j - i
		}
	}
	return res
}
