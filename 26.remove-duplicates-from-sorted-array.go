package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var k int = len(nums)
	if k == 0 || k == 1 {
		return k
	}

	var index int = 0
	for i := 1; i < k; i++ {
		if nums[index] != nums[i] {
			index = index + 1
			nums[index] = nums[i]
		}
	}

	return index + 1
}

/**
 * Accepted
 * 165/165 cases passed (8 ms)
 * Your runtime beats 87.15 % of golang submissions
 * Your memory usage beats 23 % of golang submissions (4.6 MB)
 */

// @lc code=end
