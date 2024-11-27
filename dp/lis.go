package dp

import (
	"fmt"
	"sort"
)

func RunLISTest() {
	lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
}

// Dynamic Programming with Binary Search
// Time complexity : O(nlogn). Binary search takes nlogn time and it is called n times.
// Space complexity : O(n). dp array of size n is used.
func lengthOfLIS(nums []int) int {
	// dp keeps some of the visited element in a sorted list, and its size is
	// lengthOfLIS sofar. It always keeps the our best chance to build a LIS in the future.
	dp := []int{}
	for _, num := range nums {
		fmt.Println("before search", num, dp)
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			// If num is the biggest, we should add it into the end of dp.
			dp = append(dp, num)
			fmt.Println("i == len(dp)", num, i, dp)
		} else {
			// If num is not the biggest, we should keep it in dp and replace the
			// previous num in this position. Because even if num can't build longer
			// LIS directly, it can help build a smaller dp, and we will have the
			// best chance to build a LIS in the future. All elements before this
			// position will be the best(smallest) LIS sor far.
			dp[i] = num
			fmt.Println("i <> len(dp)", num, i, dp)
		}
	}
	// dp doesn't keep LIS, and only keep the lengthOfLIS.
	return len(dp)
}
