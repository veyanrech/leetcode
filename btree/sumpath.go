package btree

func hasPathSum(root *TreeNode, targetSum int) bool {
	res := false
	sum := 0
	R := rec{
		res:    res,
		target: targetSum,
	}
	R.recur(root, sum)
	return res
}

type rec struct {
	res    bool
	target int
}

func (r *rec) recur(node *TreeNode, sumPassed int) {
	if node != nil {
		sumPassed += node.Val
	} else {
		return
	}
	if sumPassed == r.target && node.Left == nil && node.Right == nil {
		if r.res == false {
			r.res = true
			return
		}
	} else {
		if node.Left != nil {
			r.recur(node.Left, sumPassed)
		}
		if node.Right != nil {
			r.recur(node.Right, sumPassed)
		}
	}
}
