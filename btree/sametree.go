package btree

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	bl, br := false, false
	if p == nil && q == nil {
		return true
	} else {
		if p == nil || q == nil {
			return false
		}
	}
	if p.Val == q.Val {

		if p.Left == q.Left {
			bl = true
		} else {
			bl = IsSameTree(p.Left, q.Left)
		}
		if p.Right == q.Right {
			br = true
		} else {
			br = IsSameTree(p.Right, q.Right)
		}
	} else {
		return false
	}

	return bl && br
}
