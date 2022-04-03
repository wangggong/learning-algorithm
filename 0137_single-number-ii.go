/*
 * @lc app=leetcode.cn id=single-number-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode-cn.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (72.05%)
 * Total Accepted:    112.9K
 * Total Submissions: 156.6K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 1e4
 * -(1 << 31) <= nums[i] <= (1 << 31) - 1
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 */

// 位运算, 每位取除 3 的余数.
func singleNumber(nums []int) int {
	ans := 0
	for i := 32; i >= 0; i-- {
		cnt := 0
		for _, n := range nums {
			n += 1 << 31
			if (n>>i)&1 != 0 {
				cnt++
			}
		}
		ans <<= 1
		if cnt%3 != 0 {
			ans |= 1
		}
	}
	return ans - (1 << 31)
}
