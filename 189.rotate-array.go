package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)

	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}

/**
 * 37/37 cases passed (32 ms)
 * Your runtime beats 80 % of golang submissions
 * Your memory usage beats 18.21 % of golang submissions (8.3 MB)
 */
// @lc code=end
