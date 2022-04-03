/*
 * @lc app=leetcode.cn id=delete-operation-for-two-strings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [583] 两个字符串的删除操作
 *
 * https://leetcode-cn.com/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (63.22%)
 * Total Accepted:    65.8K
 * Total Submissions: 103.5K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
 *
 * 每步 可以删除任意一个字符串中的一个字符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: word1 = "sea", word2 = "eat"
 * 输出: 2
 * 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
 *
 *
 * 示例  2:
 *
 *
 * 输入：word1 = "leetcode", word2 = "etco"
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= word1.length, word2.length <= 500
 * word1 和 word2 只包含小写英文字母
 *
 *
 */

const MAXN int = 500

var dp [MAXN + 5][MAXN + 5]int

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	for i := 0; i < MAXN+5; i++ {
		for j := 0; j < MAXN+5; j++ {
			dp[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j+1] = dp[i][j]
			if word1[i] == word2[j] {
				dp[i+1][j+1]++
			}
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i+1][j])
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j+1])
		}
	}
	return n + m - 2*dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
