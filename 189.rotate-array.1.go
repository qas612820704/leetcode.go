package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}
func reverse(nums []int) {
	k := len(nums)
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
}

/**
 * 37/37 cases passed (32 ms)
 * Your runtime beats 80 % of golang submissions
 * Your memory usage beats 18.21 % of golang submissions (8.3 MB)
 */
// @lc code=end
