package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(canFinish(2, [][]int{{1, 0}}), true))
	fmt.Println("2", reflect.DeepEqual(canFinish(2, [][]int{{1, 0}, {0, 1}}), false))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// need to make a map with prerequesites in form map[course][]prereq
	// loop through map
	// call dfs
	graph := make(map[int][]int)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}
	visited := make(map[int]bool)

	var dfs func(course int) bool

	dfs = func(course int) bool {
		if _, ok := visited[course]; ok {
			return false
		}

		if len(graph[course]) == 0 {
			return true
		}

		visited[course] = true

		for _, pre := range graph[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visited, course)

		graph[course] = []int{}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
