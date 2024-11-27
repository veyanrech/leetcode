package jumpgame2

func jump(nums []int) int {

	n := len(nums)
	if n == 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = n
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}

	return dp[n-1]
}

func RunJumpGame2() int {
	nums := []int{1, 2, 3}
	return jump(nums)
}
