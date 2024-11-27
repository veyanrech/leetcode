package dp

import "fmt"

func RunCanjumpTest() {
	tc := []int{1, 2, 3, 0, 5}
	fmt.Println(canJump(tc))
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	res := false

	from := len(nums) - 2
	to := len(nums) - 1
	for i := 0; i < len(nums) && from >= 0; {
		if nums[from]+from >= to {
			i++
			res = true
			from = len(nums) - 2 - i
			to = len(nums) - 1 - i
		} else {
			i++
			res = false
			from = len(nums) - 2 - i
		}
	}

	return res
}

// func solve(startIndex int, ch map[int]int, nums []int) int {

// 	v, ok := ch[startIndex]
// 	if ok {
// 		return v
// 	}

// 	if startIndex+nums[startIndex] == len(nums)-1 {
// 		ch[startIndex] = len(nums) - 1
// 	}

// 	for i := startIndex + 1; i <= nums[startIndex]; i++ {
// 		ch[i] = solve(i, ch, nums)
// 	}

// 	return false
// }
