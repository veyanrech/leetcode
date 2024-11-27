package scalar

import (
	"fmt"
	"math"
)

func MaxDotProduct(nums1 []int, nums2 []int) int {

	var a [][]int = make([][]int, len(nums1), len(nums1))
	// var maximums [][]int = make([][]int, 0)
	// max := 0
	maxsum := math.MinInt
	// maxIndex1 := -1
	maxIndex2 := -1
	maxIndex1 := -1
	product := 0
	for i, v := range nums1 {
		a[i] = make([]int, 0, len(nums2))
		product = 0
		localmax := 0
		for i2, v2 := range nums2 {
			product = v * v2
			a[i] = append(a[i], product)

			if i2 == 0 {
				localmax = product
			}

			if i == 0 && i2 == 0 {
				maxsum = product
				localmax = product
				maxIndex2 = i2
				maxIndex1 = i
				continue
			}

			if product > maxsum && product >= localmax {
				if maxsum+product > maxsum && maxsum+product > localmax && maxIndex2 < i2 && maxIndex1 != i {
					localmax = maxsum + product
					maxIndex2 = i2
					maxIndex1 = i
				} else {
					if product > localmax && i2 < maxIndex2 {
						localmax = product
						maxIndex2 = i2
						maxIndex1 = i
					}
				}
			} else {
				if maxsum+product > maxsum && maxsum+product > localmax && maxIndex2 < i2 && maxIndex1 != i {
					localmax = maxsum + product
					maxIndex2 = i2
					maxIndex1 = i
				}
			}

		}
		maxsum = localmax
	}
	fmt.Println(maxsum, maxIndex1, maxIndex2)
	return maxsum

}
