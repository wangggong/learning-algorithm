/*
 * @lc app=leetcode.cn id=perfect-squares lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (64.32%)
 * Total Accepted:    257.6K
 * Total Submissions: 400.2K
 * Testcase Example:  '12'
 *
 * 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
 *
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：3
 * 解释：12 = 4 + 4 + 4
 *
 * 示例 2：
 *
 *
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for i := 0; i*i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i*i {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
