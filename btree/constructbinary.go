package btree

import "fmt"

func TestbuildTree() {
	a1 := []int{3, 9, 20, 15, 7}
	a2 := []int{9, 3, 15, 20, 7}
	res := buildTree(a1, a2)
	fmt.Println(res)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var root *TreeNode

	pv := peek(&preorder)
	mid := index(&inorder, pv)

	root = CreateTreeNode(pv, nil, nil)

	root.Left = buildTree(preorder[:mid], inorder[:mid])
	root.Right = buildTree(preorder[mid:], inorder[mid+1:])

	return root
}

func peek(q *[]int) int {
	var retval int
	retval = (*q)[0]
	(*q) = (*q)[1:]

	return retval
}

func index(a *[]int, val int) int {
	for i, v := range *a {
		if v == val {
			return i
		}
	}
	return 0
}

func CreateTreeNode(v int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  l,
		Right: r,
	}
}

func peekTreenode(q *[]*TreeNode) *TreeNode {
	retval := (*q)[0]
	(*q) = (*q)[1:]
	return retval
}
