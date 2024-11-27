package arrays

import "fmt"

func TestsearchInsert() {
	a := []int{1, 3}
	t := 2
	fmt.Println(searchInsert(a, t))
}

func searchInsert(nums []int, target int) int {
	middle := 0
	s := 0
	e := len(nums)
	for s <= e {
		middle = (s + e) / 2
		if middle < len(nums) {
			if target == nums[middle] {
				return middle
			}
			if target < nums[middle] {
				e = middle - 1
			}
			if target > nums[middle] {
				s = middle + 1
			}
		} else {
			break
		}
	}
	return s
}
