/*
 * @lc app=leetcode.cn id=longest-common-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1143] 最长公共子序列
 *
 * https://leetcode-cn.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (64.15%)
 * Total Accepted:    213.9K
 * Total Submissions: 333.5K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
 *
 * 一个字符串的 子序列
 * 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 *
 *
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
 *
 *
 * 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：text1 = "abcde", text2 = "ace"
 * 输出：3
 * 解释：最长公共子序列是 "ace" ，它的长度为 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc" ，它的长度为 3 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= text1.length, text2.length <= 1000
 * text1 和 text2 仅由小写英文字符组成。
 *
 *
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	b1, b2 := []byte(text1), []byte(text2)
	n, m := len(b1), len(b2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			dp[i][j] = max(dp[i][j], dp[i][j-1])
			if b1[i-1] == b2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
