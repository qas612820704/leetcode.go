package main

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var profit int = 0

	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			profit = profit + diff
		}
	}

	return profit
}

/**
 * 200/200 cases passed (4 ms)
 * Your runtime beats 94.93 % of golang submissions
 * Your memory usage beats 100 % of golang submissions (3.1 MB)
 */
// @lc code=end
