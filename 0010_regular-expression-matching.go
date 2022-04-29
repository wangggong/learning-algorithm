/*
 * @lc app=leetcode.cn id=regular-expression-matching lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode-cn.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (31.61%)
 * Total Accepted:    264.9K
 * Total Submissions: 838.2K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 *
 *
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 *
 *
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aa", p = "a"
 * 输出：false
 * 解释："a" 无法匹配 "aa" 整个字符串。
 *
 *
 * 示例 2:
 *
 *
 * 输入：s = "aa", p = "a*"
 * 输出：true
 * 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "ab", p = ".*"
 * 输出：true
 * 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * 1 <= p.length <= 30
 * s 只包含从 a-z 的小写字母。
 * p 只包含从 a-z 的小写字母，以及字符 . 和 *。
 * 保证每次出现字符 * 时，前面都匹配到有效的字符
 *
 *
 */
const maxn int = 30

var dp [maxn + 5][maxn + 5]bool

func match(text, pattern []byte, i, j int) bool {
	if i == 0 || j == 0 {
		return i == 0 && j == 0
	}
	if pattern[j-1] == '.' {
		return true
	}
	return text[i-1] == pattern[j-1]
}

func isMatch(s string, p string) bool {
	text, pattern := []byte(s), []byte(p)
	n, m := len(text), len(pattern)
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = false
		}
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		if i-1 > 0 && pattern[i-1] == '*' {
			dp[0][i] = dp[0][i] || dp[0][i-2]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j-1 > 0 && pattern[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(text, pattern, i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if match(text, pattern, i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}
		}
	}
	return dp[n][m]
}
