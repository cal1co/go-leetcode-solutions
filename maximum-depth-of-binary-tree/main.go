package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(maxDepth(t1) == 3)

}

func maxDepth(root *TreeNode) int {
	// helper func
	// recursively call helper and inc count and return max count recorded by traversal
	return countMax(root, 0)
}

func countMax(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	count++
	return max(countMax(root.Left, count), countMax(root.Right, count))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
