package removeduplicates

import "fmt"

func RemoveDuplicates(nums []int) int {

	counter := 1
	pointer := 0
	v := 0
	for i := 0; i < len(nums)-pointer; i++ {
		if i == 0 {
			continue
		}
		v = nums[i]
		if v == nums[i-1] {
			counter++
		} else {
			counter = 1
		}

		if counter > 2 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, (v+i)*(-1))
			counter--
			i--
			pointer++
		}
		fmt.Println(i, nums)
	}
	// fmt.Println(nums[:len(nums)-pointer])
	return len(nums[:len(nums)-pointer])
}
