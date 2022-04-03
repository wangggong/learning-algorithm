/*
 * @lc app=leetcode.cn id=factorial-trailing-zeroes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [172] 阶乘后的零
 *
 * https://leetcode-cn.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Medium (45.28%)
 * Total Accepted:    116.7K
 * Total Submissions: 250.4K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n ，返回 n! 结果中尾随零的数量。
 *
 * 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：0
 * 解释：3! = 6 ，不含尾随 0
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 5
 * 输出：1
 * 解释：5! = 120 ，有一个尾随 0
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 10^4
 *
 *
 *
 *
 * 进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
 *
 */
func trailingZeroes(n int) int {
	ans := 0
	cur := 5
	for n >= cur {
		ans += n / cur
		cur *= 5
	}
	return ans
}
