package btree

func RunLevelsTest() {

}

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
}
*/

type levelsCalculator struct {
	levels [][]int
}

func averageOfLevels(root *TreeNode) []float64 {
	lcv := &levelsCalculator{
		levels: [][]int{},
	}
	calc(root, lcv, 0)

	averages := []float64{}

	for _, v := range lcv.levels {
		var av float64
		var sum float64 = 0
		var counter int
		for i, v2 := range v {
			sum += float64(v2)
			counter = i
		}
		av = sum / float64(counter+1)
		averages = append(averages, av)

	}

	return averages
}

func calc(root *TreeNode, lc *levelsCalculator, level int) {
	if root == nil {
		return
	}
	if lc.levels == nil {
		lc.levels = make([][]int, 0)
	}
	if len(lc.levels)-1 < level {
		lc.levels = append(lc.levels, make([]int, 0))
	}
	lc.levels[level] = append(lc.levels[level], root.Val)

	level += 1
	if root.Left != nil {
		calc(root.Left, lc, level)
	}
	if root.Right != nil {
		calc(root.Right, lc, level)
	}
}
