/*
 * @lc app=leetcode.cn id=count-numbers-with-unique-digits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [357] 统计各位数字都不同的数字个数
 *
 * https://leetcode-cn.com/problems/count-numbers-with-unique-digits/description/
 *
 * algorithms
 * Medium (52.35%)
 * Total Accepted:    38.4K
 * Total Submissions: 68.6K
 * Testcase Example:  '2'
 *
 * 给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10^n^ 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：91
 * 解释：答案应为除去 11、22、33、44、55、66、77、88、99 外，在 0 ≤ x < 100 范围内的所有数字。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：1
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 8
 *
 *
 */
func countNumbersWithUniqueDigits(n int) int {
	ans := 0
	// 傻逼! 各位数不同的 `n` 位数是可以乘法原理直接算的!
	for i := 0; i <= n; i++ {
		cur := 1
		k := 9
		for j := i - 1; j > 0; j-- {
			cur *= k
			k--
		}
		if i > 0 {
			cur *= 9
		}
		ans += cur
	}
	return ans
}
