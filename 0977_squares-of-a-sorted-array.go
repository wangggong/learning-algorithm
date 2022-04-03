/*
 * @lc app=leetcode.cn id=squares-of-a-sorted-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [977] 有序数组的平方
 *
 * https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (69.60%)
 * Total Accepted:    278.8K
 * Total Submissions: 403.1K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 解释：平方后，数组变为 [16,1,0,9,100]
 * 排序后，数组变为 [0,1,9,16,100]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1e4
 * -1e4 <= nums[i] <= 1e4
 * nums 已按 非递减顺序 排序
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 请你设计时间复杂度为 O(n) 的算法解决本问题
 *
 *
 */
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	p, q := 0, len(nums)-1
	k := n-1
	for p <= q {
		if nums[p] * nums[p] < nums[q] * nums[q] {
			ans[k] = nums[q] * nums[q]
			k--
			q--
		} else {
			ans[k] = nums[p] * nums[p]
			k--
			p++
		}
	}
	return ans
}
