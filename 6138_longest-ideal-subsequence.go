/*
 * @lc app=leetcode.cn id=longest-ideal-subsequence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6138] 最长理想子序列
 *
 * https://leetcode.cn/problems/longest-ideal-subsequence/description/
 *
 * algorithms
 * Medium (37.99%)
 * Total Accepted:    4.1K
 * Total Submissions: 10.9K
 * Testcase Example:  '"acfgbd"\n2'
 *
 * 给你一个由小写字母组成的字符串 s ，和一个整数 k 。如果满足下述条件，则可以将字符串 t 视作是 理想字符串 ：
 *
 *
 * t 是字符串 s 的一个子序列。
 * t 中每两个 相邻 字母在字母表中位次的绝对差值小于或等于 k 。
 *
 *
 * 返回 最长 理想字符串的长度。
 *
 * 字符串的子序列同样是一个字符串，并且子序列还满足：可以经由其他字符串删除某些字符（也可以不删除）但不改变剩余字符的顺序得到。
 *
 * 注意：字母表顺序不会循环。例如，'a' 和 'z' 在字母表中位次的绝对差值是 25 ，而不是 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "acfgbd", k = 2
 * 输出：4
 * 解释：最长理想字符串是 "acbd" 。该字符串长度为 4 ，所以返回 4 。
 * 注意 "acfgbd" 不是理想字符串，因为 'c' 和 'f' 的字母表位次差值为 3 。
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd", k = 3
 * 输出：4
 * 解释：最长理想字符串是 "abcd" ，该字符串长度为 4 ，所以返回 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * 0 <= k <= 25
 * s 由小写英文字母组成
 *
 *
 */
func longestIdealString(s string, k int) (ans int) {
	dp := make([]int, 'z'-'a'+1)
	for _, b := range s {
		last := make([]int, 'z'-'a'+1)
		copy(last, dp)
		for j := max(0, int(b-'a')-k); j <= min(int('z'-'a'), int(b-'a')+k); j++ {
			dp[int(b-'a')] = max(dp[int(b-'a')], last[j]+1)
		}
		// fmt.Println(dp)
	}
	for _, d := range dp {
		ans = max(ans, d)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
