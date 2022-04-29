/*
 * @lc app=leetcode.cn id=global-and-local-inversions lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [775] 全局倒置与局部倒置
 *
 * https://leetcode-cn.com/problems/global-and-local-inversions/description/
 *
 * algorithms
 * Medium (45.83%)
 * Total Accepted:    5.9K
 * Total Submissions: 12.9K
 * Testcase Example:  '[1,0,2]'
 *
 * 给你一个长度为 n 的整数数组 nums ，表示由范围 [0, n - 1] 内所有整数组成的一个排列。
 *
 * 全局倒置 的数目等于满足下述条件不同下标对 (i, j) 的数目：
 *
 *
 * 0
 * nums[i] > nums[j]
 *
 *
 * 局部倒置 的数目等于满足下述条件的下标 i 的数目：
 *
 *
 * 0
 * nums[i] > nums[i + 1]
 *
 *
 * 当数组 nums 中 全局倒置 的数量等于 局部倒置 的数量时，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,2]
 * 输出：true
 * 解释：有 1 个全局倒置，和 1 个局部倒置。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,0]
 * 输出：false
 * 解释：有 2 个全局倒置，和 1 个局部倒置。
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1
 * 0
 * nums 中的所有整数 互不相同
 * nums 是范围 [0, n - 1] 内所有数字组成的一个排列
 *
 *
 */
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	maxVal := nums[0]
	for i := 2; i < n; i++ {
		if maxVal > nums[i] {
			return false
		}
		if maxVal < nums[i-1] {
			maxVal = nums[i-1]
		}
	}
	return true
}
