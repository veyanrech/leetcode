package btree

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 */
func MaxDepth(root *TreeNode) int {
	l, r := 1, 1

	if root.Left == nil && root.Right == nil {
		return l
	}

	if root.Left != nil {
		l += MaxDepth(root.Left)
	}

	if root.Right != nil {
		r += MaxDepth(root.Right)
	}

	if l > r {
		return l
	}
	return r
}
