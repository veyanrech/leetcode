package removeelement

import "fmt"

func RemoveElement(nums []int, val int) int {

	fmt.Println(nums)

	k := 0
	j := 0
	i := 0
	for i < len(nums)-j {
		v := nums[i]
		if v == val {
			nums = append(nums[:i], nums[i+1:]...)
			// fmt.Println("1:", i, v, nums)
			nums = append(nums, v)
			// fmt.Println("2:", i, v, nums)
			j++
		} else {
			k++
			i++
		}
	}
	// fmt.Println(nums)
	return k
}
