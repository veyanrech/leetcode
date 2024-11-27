package arrays

import "sort"

func TestTriplet() {
	t := []int{3, 0, -2, -1, 1, 2, 2, 31, 1, 1, 1, -31}
	//-2 -1 0 1 2 3
	threeSum(t)
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sum := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}

		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {

				if i != j && i != k && j != k {
					if (nums[i] + nums[j] + nums[k]) == 0 {
						if len(sum) > 0 {
							if nums[k] != sum[len(sum)-1][2] || nums[j] != sum[len(sum)-1][1] || nums[i] != sum[len(sum)-1][0] {
								sum = append(sum, []int{nums[i], nums[j], nums[k]})
							}
						} else {
							sum = append(sum, []int{nums[i], nums[j], nums[k]})
						}
					} else {
						if (nums[i] + nums[j] + nums[k]) > 0 {
							break
						} else {
							continue
						}
					}
				}

				// if (nums[i] + nums[j] + nums[k]) == 0 {
				// 	if len(sum) > 0 {
				// 		if nums[k] != sum[len(sum)-1][2] || nums[j] != sum[len(sum)-1][1] || nums[i] != sum[len(sum)-1][0] {
				// 			sum = append(sum, []int{nums[i], nums[j], nums[k]})
				// 		}
				// 	} else {
				// 		sum = append(sum, []int{nums[i], nums[j], nums[k]})
				// 	}

				// }
			}
		}
	}
	return sum
}
