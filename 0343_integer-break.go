/*
 * @lc app=leetcode.cn id=integer-break lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [343] 整数拆分
 *
 * https://leetcode-cn.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (61.45%)
 * Total Accepted:    147.1K
 * Total Submissions: 239.3K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
 *
 * 返回 你可以获得的最大乘积 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 *
 * 输入: n = 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= n <= 58
 *
 *
 */

/*
 * // DP
 * func integerBreak(n int) int {
 * 	dp := make([]int, n+1)
 * 	dp[1] = 1
 * 	for i := 2; i <= n; i++ {
 * 		for j := 1; j < i; j++ {
 * 			dp[i] = max(dp[i], j*(i-j))
 * 			dp[i] = max(dp[i], dp[j]*(i-j))
 * 			// dp[i] = max(dp[i], dp[j]*dp[i-j])	// 其实不需要这个了
 * 		}
 * 	}
 * 	return dp[n]
 * }
 */

// greedy ---- 可恶, 被锋哥秀到了
//
// 大致就是有 `3` 拆 `3` 没 `3` 拆 `4`, 证明留作练习.
func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
	}
	ans *= n
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
