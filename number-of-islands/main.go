package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}) == 1)
	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}) == 3)
}
func numIslands(grid [][]byte) int {
	// loop through the grid
	// if we see land, we run dfs, marking land as visited
	// increment output
	var out int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				markLand(grid, i, j)
				out++
			}
		}
	}

	return out
}

func markLand(grid [][]byte, i, j int) {
	// conditions to check if out of bounds
	if i > len(grid)-1 || i < 0 || j > len(grid[0])-1 || j < 0 || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	markLand(grid, i+1, j)
	markLand(grid, i-1, j)
	markLand(grid, i, j+1)
	markLand(grid, i, j-1)
}
