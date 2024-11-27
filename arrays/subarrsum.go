package arrays

import (
	"fmt"
	"math"
)

func RunTestmaxSubArray() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt
	for _, v := range nums {
		if v > sum+v {
			sum = v
		} else {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
