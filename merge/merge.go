package merge

import "fmt"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for len(nums2) > 0 {
		n2 := nums2[0]
		n1 := nums1[i]

		if nums1[i] <= nums2[0] && (len(nums1) > i+1 && (nums2[0] <= nums1[i+1] || nums1[i+1] == 0)) {
			nums1 = append(nums1[:i+1], nums1[i:(m+n-1)]...)
			nums1[i+1] = n2
			nums2 = nums2[1:]
		}

		if n1 > n2 {
			nums2[0] = n1
			nums1[i] = n2
		}

		i++
	}
}

func Merge2(nums1 []int, m int, nums2 []int, n int) {

	// if m == 0 {
	// 	nums1 = append([]int{}, nums2[:n]...)
	// 	fmt.Println(len(nums1), cap(nums1), nums1)
	// 	return
	// }

	if n == 0 {
		fmt.Println(nums1)
		return
	}

	i := 0
	j := 0

	mn := m + n

	for i < mn {

		if j >= len(nums2) {
			j = 0
			i = 0
		}

		n1 := nums1[i]
		n2 := nums2[j]

		if n1 > n2 {
			nums1[i] = n2
			nums2[j] = n1
			i = 0
			j++
			continue
		}

		if n1 == 0 && i >= m {
			nums1[i] = n2
			j++
			i++
			continue
		}

		if n1 < n2 {

			if nums1[i+1] == 0 {
				i++
				continue
			}

			i++
		}

		if n1 == n2 {
			nums1 = append(nums1[:i+1], nums1[i:(mn-1)]...)
			nums1[i+1] = n2
			j++
			i++
		}
	}

	fmt.Println(nums1)
}

func Merge3(nums1 []int, m int, nums2 []int, n int) {

	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

}
