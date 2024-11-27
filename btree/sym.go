package btree

func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	qu := []*TreeNode{}

	qu = append(qu, root)
	qu = append(qu, root)

	for len(qu) > 0 {
		l, r := pop(&qu), pop(&qu)
		if l == r && l == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		qu = append(qu, l.Left)
		qu = append(qu, r.Right)
		if l != r {
			qu = append(qu, l.Right)
			qu = append(qu, r.Left)
		}
	}

	return true
}

func pop(q *[]*TreeNode) *TreeNode {
	retval := (*q)[len((*q))-1]
	(*q) = (*q)[:len((*q))-1]
	return retval
}
