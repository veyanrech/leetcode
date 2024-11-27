package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return []int{}
	} else {
		arr = append(arr, root.Val)
		arr = append(arr, dif(root.Left, root.Right)...)
	}
	return arr
}

func dif(l, r *TreeNode) []int {
	retval := []int{}
	if l == nil && r != nil {
		retval = append(retval, r.Val)
		retval = append(retval, dif(r.Left, r.Right)...)
		return retval
	}
	if l != nil && r == nil {
		retval = append(retval, l.Val)
		retval = append(retval, dif(l.Left, l.Right)...)
		return retval
	}
	if l == nil && r == nil {
		return []int{}
	}
	if l.Val > r.Val {
		retval = append(retval, l.Val)
		retval = append(retval, dif(l.Left, l.Right)...)
		retval = append(retval, dif(r.Left, r.Right)...)
		return retval
	}
	retval = append(retval, r.Val)
	retval = append(retval, dif(l.Left, l.Right)...)
	retval = append(retval, dif(r.Left, r.Right)...)
	return retval

}
