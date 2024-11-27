package btree

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type allvals struct {
	levels []int
}

func getMinimumDifference(root *TreeNode) int {
	av := &allvals{
		levels: make([]int, 0),
	}

	appendVal(root, av)

	sort.Slice(av.levels, func(i, j int) bool {
		return av.levels[j] < av.levels[i]
	})

	var res int

	if len(av.levels) > 1 {
		res = abs(av.levels[1] - av.levels[0])
		for i := 2; i < len(av.levels); i++ {
			if res < abs(av.levels[i]-av.levels[i-1]) {
				res = abs(av.levels[i] - av.levels[i-1])
			}
		}
	}

	return res
}

func appendVal(root *TreeNode, nodevals *allvals) {
	if root != nil {
		nodevals.levels = append(nodevals.levels, root.Val)
		if root.Left != nil {
			appendVal(root.Left, nodevals)
		}
		if root.Right != nil {
			appendVal(root.Right, nodevals)
		}
	}
}

func abs(ival int) int {
	if ival < 0 {
		return ival * -1
	}
	return ival
}
