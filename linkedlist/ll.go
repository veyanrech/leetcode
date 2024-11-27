package linkedlist

type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) add(val int) {
	newNode := &ListNode{Val: val}
	if ll.head == nil {
		ll.head = newNode
	} else {
		cur := ll.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	realLeft := left - 1
	realRight := right - 1
	var sliceFromList []int = make([]int, 0)
	lnode := head
	if lnode != nil {
		sliceFromList = append(sliceFromList, lnode.Val)
		for lnode.Next != nil {
			lnode = lnode.Next
			sliceFromList = append(sliceFromList, lnode.Val)
		}
	}
	newList := LinkedList{}
	for i := 0; i < left; i++ {
		newList.add(sliceFromList[i])
	}

	for i := realRight; i >= realLeft; i-- {
		newList.add(sliceFromList[i])
	}

	for i := realRight; i < len(sliceFromList); i++ {
		newList.add(sliceFromList[i])
	}

	return newList.head
}
