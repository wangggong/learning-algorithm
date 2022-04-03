/*
 * @lc app=leetcode.cn id=minimum-ascii-delete-sum-for-two-strings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [712] 两个字符串的最小ASCII删除和
 *
 * https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/description/
 *
 * algorithms
 * Medium (67.70%)
 * Total Accepted:    19K
 * Total Submissions: 28K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s1 = "sea", s2 = "eat"
 * 输出: 231
 * 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
 * 在 "eat" 中删除 "t" 并将 116 加入总和。
 * 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s1 = "delete", s2 = "leet"
 * 输出: 403
 * 解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
 * 将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
 * 结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
 * 如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= s1.length, s2.length <= 1000
 * s1 和 s2 由小写英文字母组成
 *
 *
 */

const MAXN int = 1000

var dp [MAXN + 5][MAXN + 5]int

func minimumDeleteSum(s1 string, s2 string) int {
	b1, b2 := []byte(s1), []byte(s2)
	for i := 0; i < MAXN+5; i++ {
		for j := 0; j < MAXN+5; j++ {
			dp[i][j] = MAXN * 256
		}
	}
	n, m := len(b1), len(b2)
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + int(b1[i-1])
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = dp[0][i-1] + int(b2[i-1])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if b1[i-1] == b2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+int(b2[j-1]))
				dp[i][j] = min(dp[i][j], dp[i-1][j]+int(b1[i-1]))
			}
		}
	}
	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
