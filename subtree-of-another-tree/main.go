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
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	t2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(t1, t2))

}

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isEquals(root, subRoot) || isEquals(root.Left, subRoot) || isEquals(root.Right, subRoot)
}

func isEquals(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isEquals(p.Right, q.Right) && isEquals(p.Left, q.Left)
}
