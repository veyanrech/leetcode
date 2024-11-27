package duplelement

func DuplElement(nums []int) int {

	// fmt.Println(nums)

	k := 0
	i := 1
	for i < len(nums)-k {
		v := nums[i]
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, v)
			// fmt.Println(i, "a:", nums)
			k++
			i--
		}
		i++
	}

	// fmt.Println(nums)

	return len(nums) - k
}
