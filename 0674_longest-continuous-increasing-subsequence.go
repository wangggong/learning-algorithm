/*
 * @lc app=leetcode.cn id=longest-continuous-increasing-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [674] 最长连续递增序列
 *
 * https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (51.76%)
 * Total Accepted:    115.2K
 * Total Submissions: 222.5K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
 *
 * 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l ，都有 nums[i] < nums[i + 1] ，那么子序列
 * [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,5,4,7]
 * 输出：3
 * 解释：最长连续递增序列是 [1,3,5], 长度为3。
 * 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2]
 * 输出：1
 * 解释：最长连续递增序列是 [2], 长度为1。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1e4
 * -1e9 <= nums[i] <= 1e9
 *
 *
 */

import "math"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	// assert n > 0;
	ans := 0
	last := math.MinInt32
	curr := 0
	for i := 0; i < n; i++ {
		if last < nums[i] {
			curr++
		} else {
			curr = 1
		}
		last = nums[i]
		ans = max(ans, curr)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
