/*
 * @lc app=leetcode.cn id=distinct-subsequences lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [115] 不同的子序列
 *
 * https://leetcode-cn.com/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (51.86%)
 * Total Accepted:    79.2K
 * Total Submissions: 152.7K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
 *
 * 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE"
 * 的一个子序列，而 "AEC" 不是）
 *
 * 题目数据保证答案符合 32 位带符号整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * rabbbit
 * rabbbit
 * rabbbit
 *
 * 示例 2：
 *
 *
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length, t.length <= 1000
 * s 和 t 由英文字母组成
 *
 *
 */
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// 初始状态: 在任何字符串中匹配 "" (空串) 都只能匹配一次.
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 如果不依赖当前字符匹配, `s[0..i]` 与 `t[0..j]` 匹配的种数等价于 `s[0..i-1]` 与 `t[0..j]` 的匹配种数;
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				// 如果当前字符能匹配 (`s[i-1] == t[j-1]`), 支持按当前字符匹配, 结果加上 `s[0..i-1]` 与 `t[0..j-1]` 的匹配种数;
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}
