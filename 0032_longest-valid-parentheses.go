/*
 * @lc app=leetcode.cn id=longest-valid-parentheses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (35.92%)
 * Total Accepted:    225.7K
 * Total Submissions: 627.7K
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 3 * 1e4
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */
func longestValidParentheses(s string) int {
	bytes := []byte(s)
	n := len(bytes)
	var S []int
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n + 1
	}
	ans := 0
	for i, b := range bytes {
		if b == '(' {
			S = append(S, i)
			continue
		}
		if len(S) == 0 {
			continue
		}
		left := S[len(S)-1]
		S = S[:len(S)-1]
		dp[i] = min(dp[i], left)
		for left > 0 && left != n+1 && dp[left-1] < dp[i] {
			dp[i] = dp[left-1]
			left = dp[left-1]
		}
		ans = max(ans, i-dp[i]+1)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
