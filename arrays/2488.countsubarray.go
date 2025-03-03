package arrays

func countSubarrays(nums []int, k int) int {

	n := len(nums)
	count := 0

	kindex := -1
	for i := 0; i < n; i++ {
		if nums[i] == k {
			kindex = i
			break
		}
	}

	if kindex == -1 {
		return 0
	}

	count++

	balanceFrequency := make(map[int]int)
	balanceFrequency[0] = 1

	balance := 0
	for i := kindex + 1; i < n; i++ {
		if nums[i] > k {
			balance++
		} else if nums[i] < k {
			balance--
		}

		balanceFrequency[balance]++
	}

	return count

}
