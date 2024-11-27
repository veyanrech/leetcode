package btree

import "fmt"

func TestPostorder() {
	r := buildTree2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	rootVal := pop2(&postorder)
	indexRoot := findIndex(rootVal, inorder)

	root := createTreeNode(rootVal, nil, nil)

	if len(postorder) == 0 {
		return root
	} else {
		root.Left = buildTree2(inorder[:indexRoot], postorder[:indexRoot])
		root.Right = buildTree2(inorder[indexRoot+1:], postorder[indexRoot:])
	}

	return root
}

func createTreeNode(v int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  l,
		Right: r,
	}
}

func pop2(ints *[]int) int {
	temp := (*ints)[len((*ints))-1]
	(*ints) = (*ints)[:len((*ints))-1]
	return temp
}

func findIndex(val int, ints []int) int {
	var i, v int
	for i, v = range ints {
		if v == val {
			break
		}
	}
	return i
}
