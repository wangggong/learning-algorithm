/*
 * @lc app=leetcode.cn id=find-three-consecutive-integers-that-sum-to-a-given-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5997] 找到和为给定整数的三个连续整数
 *
 * https://leetcode-cn.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/description/
 *
 * algorithms
 * Medium (66.12%)
 * Total Accepted:    3.5K
 * Total Submissions: 5.3K
 * Testcase Example:  '33'
 *
 * 给你一个整数 num ，请你返回三个连续的整数，它们的 和 为 num 。如果 num 无法被表示成三个连续整数的和，请你返回一个 空 数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：num = 33
 * 输出：[10,11,12]
 * 解释：33 可以表示为 10 + 11 + 12 = 33 。
 * 10, 11, 12 是 3 个连续整数，所以返回 [10, 11, 12] 。
 *
 *
 * 示例 2：
 *
 * 输入：num = 4
 * 输出：[]
 * 解释：没有办法将 4 表示成 3 个连续整数的和。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= num <= 10^15
 *
 *
 */
func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	p := num / 3
	return []int64{p - 1, p, p + 1}
}
