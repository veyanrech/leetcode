package dp

import "fmt"

func RunScTest() {
	tc := []int{7, 1, 5, 4, 5, 4, 15}
	fmt.Println(maxProfit2(tc))
}

// var cch map[[2]int]int = map[[2]int]int{}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxsum := 0

	for i := len(prices) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			sell := prices[i] - prices[j]
			if sell > 0 {
				maxsum = max(sell+maxProfit(prices[:j]), maxsum)
				// cch[[2]int{i, j}] = maxsum
			}
		}
	}

	return maxsum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {
	profit := 0
	// no profit available at index 0, so start at index 1
	for i := 1; i < len(prices); i++ {
		// Only when the current price is higher than the previous
		// can we make a profitable sale.
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
		// The else case is handled implicitly by incrementing i.
		// When prices[i] is lower than prices[i-1], we just skip
		// the sale.
		//
		// If this is confusing, try keeping a `min` variable yourself,
		// and you'll soon realise it's redundant.
	}
	return profit
}
